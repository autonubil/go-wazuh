package ossec

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"testing"
	"time"

	"go.uber.org/zap"
)

// unregisteredRawEvent stands in for the integration events the wodles put into
// QueuePosting.Raw. Nothing registers it with gob up front, which is exactly the
// situation the late registration has to cope with.
type unregisteredRawEvent struct {
	IntegrationEvent
	Payload map[string]any
}

func gobEncode(v any) error {
	var buf bytes.Buffer
	return gob.NewEncoder(&buf).Encode(v)
}

// TestRegisterGobTypeNormalisesPointerAndValue guards the panic that took the
// remotelog wodle down: gob derives a different name for a pointer than for its
// element, and registering both forms of one type is fatal.
func TestRegisterGobTypeNormalisesPointerAndValue(t *testing.T) {
	type ptrFirst struct{ A string }
	type valueFirst struct{ B string }

	// A pointer must be normalised to the value form ...
	if _, registered := registerGobType(&ptrFirst{}); !registered {
		t.Error("registerGobType(&ptrFirst{}) did not register the type")
	}
	// ... so that the value form afterwards is a no-op rather than a panic.
	if _, registered := registerGobType(ptrFirst{}); registered {
		t.Error("registerGobType(ptrFirst{}) registered the same type twice")
	}

	if _, registered := registerGobType(valueFirst{}); !registered {
		t.Error("registerGobType(valueFirst{}) did not register the type")
	}
	if _, registered := registerGobType(&valueFirst{}); registered {
		t.Error("registerGobType(&valueFirst{}) registered the same type twice")
	}

	// QueuePosting is registered by init() in its value form. Handing the
	// pointer to it is what used to panic with
	// "registering duplicate names for *ossec.QueuePosting".
	if _, registered := registerGobType(&QueuePosting{}); registered {
		t.Error("registerGobType(&QueuePosting{}) re-registered a type init() already knows")
	}

	// Registering the value form covers both values and pointers in an interface.
	if err := gobEncode(&QueuePosting{Raw: ptrFirst{"x"}}); err != nil {
		t.Errorf("encode value in Raw: %v", err)
	}
	if err := gobEncode(&QueuePosting{Raw: &ptrFirst{"x"}}); err != nil {
		t.Errorf("encode pointer in Raw: %v", err)
	}
}

func TestRegisterGobTypeIgnoresNil(t *testing.T) {
	if typ, registered := registerGobType(nil); typ != nil || registered {
		t.Errorf("registerGobType(nil) = (%v, %v), want (nil, false)", typ, registered)
	}
	var nilRaw any
	if _, registered := registerGobType(nilRaw); registered {
		t.Error("registerGobType of a nil interface registered something")
	}
}

// TestInitRegistersJSONContainerTypes covers the arbitrary JSON some wodles hand
// on: a decoded object carries []any and map[string]any inside interfaces.
func TestInitRegistersJSONContainerTypes(t *testing.T) {
	var blob map[string]any
	if err := json.Unmarshal([]byte(`{"s":"x","n":1.5,"b":true,"arr":[1,"two",{"k":"v"}]}`), &blob); err != nil {
		t.Fatal(err)
	}
	if err := gobEncode(&QueuePosting{Raw: blob}); err != nil {
		t.Errorf("encode decoded JSON in Raw: %v", err)
	}
}

// TestOpenQueueEnqueuesUnregisteredRawType is the regression test for the crash:
// a posting whose Raw holds a type gob has never seen must end up in the queue
// instead of panicking the process from the enqueue goroutine.
func TestOpenQueueEnqueuesUnregisteredRawType(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := &Client{
		AgentKey: &AgentKey{AgentID: "000"},
		basePath: t.TempDir(),
		logger:   zap.NewNop(),
	}

	input, q, err := a.openQueue(ctx)
	if err != nil {
		t.Fatalf("openQueue: %v", err)
	}
	defer q.Close()

	input <- &QueuePosting{
		TargetQueue: LOCALFILE_MQ,
		Location:    "test",
		ProgramName: "test",
		Raw: unregisteredRawEvent{
			IntegrationEvent: IntegrationEvent{Integration: "test"},
			Payload:          map[string]any{"list": []any{1, "two"}},
		},
	}

	deadline := time.After(10 * time.Second)
	for q.Size() == 0 {
		select {
		case <-deadline:
			t.Fatal("posting with an unregistered Raw type never reached the queue")
		case <-time.After(10 * time.Millisecond):
		}
	}

	item, err := q.Dequeue()
	if err != nil {
		t.Fatalf("dequeue: %v", err)
	}
	posting, ok := item.(*QueuePosting)
	if !ok {
		t.Fatalf("dequeued %T, want *QueuePosting", item)
	}
	if _, ok := posting.Raw.(unregisteredRawEvent); !ok {
		t.Errorf("Raw round-tripped as %T, want unregisteredRawEvent", posting.Raw)
	}
	if posting.Timestamp.IsZero() {
		t.Error("enqueue did not stamp a timestamp")
	}
}

// TestOpenQueueSkipsNilPosting makes sure a nil item is dropped rather than
// dereferenced - the nil check used to sit after the first field access.
func TestOpenQueueSkipsNilPosting(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := &Client{AgentKey: &AgentKey{AgentID: "000"}, basePath: t.TempDir(), logger: zap.NewNop()}
	input, q, err := a.openQueue(ctx)
	if err != nil {
		t.Fatalf("openQueue: %v", err)
	}
	defer q.Close()

	input <- nil
	input <- &QueuePosting{Location: "test", ProgramName: "test", Raw: "after the nil"}

	deadline := time.After(10 * time.Second)
	for q.Size() == 0 {
		select {
		case <-deadline:
			t.Fatal("enqueue goroutine stopped after a nil posting")
		case <-time.After(10 * time.Millisecond):
		}
	}
	if size := q.Size(); size != 1 {
		t.Errorf("queue size = %d, want 1 (the nil posting must be dropped)", size)
	}
}
