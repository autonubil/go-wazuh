package wazuh

import (
	"fmt"
	"strconv"
	"strings"
)

type StrUInt uint64

func (t StrUInt) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(t), 10)), nil
}

func (t *StrUInt) UnmarshalJSON(s []byte) (err error) {
	r := strings.Replace(string(s), `"`, ``, -1)
	q, err := strconv.ParseUint(r, 10, 64)
	if err == nil {
		*(*uint64)(t) = q
		return
	}

	return fmt.Errorf("failed to unmarshal %s [%s] (%v)", r, string(s), err.Error())
}

type GeoPoint struct {
	Lat  float64 `json:"lat,omitempty" opensearch:"float"`
	Long float64 `json:"long,omitempty" opensearch:"float"`
}

type Manager struct {
	Name *string `json:"name,omitempty" opensearch:"keyword"`
}

type Cluster struct {
	Name *string `json:"name,omitempty" opensearch:"keyword"`
	Node *string `json:"node,omitempty" opensearch:"keyword"`
}

type EffectiveUser struct {
	ID   *string `json:"id,omitempty" opensearch:"keyword"`
	Name *string `json:"name,omitempty" opensearch:"keyword"`
}

type Group struct {
	ID   *string `json:"id,omitempty" opensearch:"keyword"`
	Name *string `json:"name,omitempty" opensearch:"keyword"`
}

type LoginUser struct {
	ID   *string `json:"id,omitempty" opensearch:"keyword"`
	Name *string `json:"name,omitempty" opensearch:"keyword"`
}

type SyscheckAuditProcess struct {
	ID   *string `json:"id,omitempty" opensearch:"keyword"`
	Name *string `json:"name,omitempty" opensearch:"keyword"`
	Ppid *string `json:"ppid,omitempty" opensearch:"keyword"`
}

type User struct {
	ID   *string `json:"id,omitempty" opensearch:"keyword"`
	Name *string `json:"name,omitempty" opensearch:"keyword"`
}

type SyscheckAudit struct {
	EffectiveUser *EffectiveUser        `json:"effective_user,omitempty" opensearch:"object"`
	Group         *Group                `json:"group,omitempty" opensearch:"object"`
	LoginUser     *LoginUser            `json:"login_user,omitempty" opensearch:"object"`
	Process       *SyscheckAuditProcess `json:"syscheck_audit_process,omitempty" opensearch:"object"`
	User          *User                 `json:"user,omitempty" opensearch:"object"`
}

type Syscheck struct {
	UidAfter     *StrUInt       `json:"uid_after,omitempty" opensearch:"long"`
	Md5Before    *string        `json:"md5_before,omitempty" opensearch:"keyword"`
	MtimeAfter   *int64         `json:"mtime_after,omitempty" opensearch:"long"`
	GidAfter     *StrUInt       `json:"gid_after,omitempty" opensearch:"long"`
	Diff         *string        `json:"diff,omitempty" opensearch:"keyword"`
	Path         *string        `json:"path,omitempty" opensearch:"keyword"`
	PermAfter    *string        `json:"perm_after,omitempty" opensearch:"keyword"`
	InodeAfter   *string        `json:"inode_after,omitempty" opensearch:"keyword"`
	InodeBefore  *string        `json:"inode_before,omitempty" opensearch:"keyword"`
	MtimeBefore  *int64         `json:"mtime_before,omitempty" opensearch:"keyword"`
	UnameBefore  *string        `json:"uname_before,omitempty" opensearch:"keyword"`
	Sha256After  *string        `json:"sha256_after,omitempty" opensearch:"keyword"`
	HardLinks    *string        `json:"hard_links,omitempty" opensearch:"keyword"`
	UidBefore    *StrUInt       `json:"uid_before,omitempty" opensearch:"long"`
	Event        *string        `json:"event,omitempty" opensearch:"keyword"`
	Sha256Before *string        `json:"sha256_before,omitempty" opensearch:"keyword"`
	Sha1After    *string        `json:"sha1_after,omitempty" opensearch:"keyword"`
	GnameBefore  *string        `json:"gname_before,omitempty" opensearch:"keyword"`
	Tags         *string        `json:"tags,omitempty" opensearch:"keyword"`
	Sha1Before   *string        `json:"sha1_before,omitempty" opensearch:"keyword"`
	PermBefore   *string        `json:"perm_before,omitempty" opensearch:"keyword"`
	GnameAfter   *string        `json:"gname_after,omitempty" opensearch:"keyword"`
	UnameAfter   *string        `json:"uname_after,omitempty" opensearch:"keyword"`
	SizeAfter    *StrUInt       `json:"size_after,omitempty" opensearch:"long"`
	Mode         *string        `json:"mode,omitempty" opensearch:"keyword"`
	GidBefore    *StrUInt       `json:"gid_before,omitempty" opensearch:"long"`
	Md5After     *string        `json:"md5_after,omitempty" opensearch:"keyword"`
	SizeBefore   *StrUInt       `json:"size_before,omitempty" opensearch:"long"`
	Audit        *SyscheckAudit `json:"syscheck_audit,omitempty" opensearch:"object"`
}

type Decoder struct {
	Accumulate *uint64 `json:"accumulate,omitempty" opensearch:"long"`
	Parent     *string `json:"parent,omitempty" opensearch:"keyword"`
	Name       *string `json:"name,omitempty" opensearch:"keyword"`
	Ftscomment *string `json:"ftscomment,omitempty" opensearch:"keyword"`
	Fts        *uint64 `json:"fts,omitempty" opensearch:"long"`
}

type Agent struct {
	IP   *string `json:"ip,omitempty" opensearch:"keyword"`
	ID   *string `json:"id,omitempty" opensearch:"keyword"`
	Name *string `json:"name,omitempty" opensearch:"keyword"`
}

type Input struct {
	Type *string `json:"type,omitempty" opensearch:"keyword"`
}

type Predecoder struct {
	ProgramName *string `json:"program_name,omitempty" opensearch:"keyword"`
	Timestamp   *string `json:"timestamp,omitempty" opensearch:"keyword"`
	Hostname    *string `json:"hostname,omitempty" opensearch:"keyword"`
}

type GeoLocation struct {
	RegionName     *string   `json:"region_name,omitempty" opensearch:"keyword"`
	Latitude       *float64  `json:"latitude,omitempty" opensearch:"float"`
	RealRegionName *string   `json:"real_region_name,omitempty" opensearch:"keyword"`
	DmaCode        *uint64   `json:"dma_code,omitempty" opensearch:"long"`
	Location       *GeoPoint `json:"location,omitempty" opensearch:"geo_point"`
	Coordinates    *float64  `json:"coordinates,omitempty" opensearch:"float"`
	CountryCode3   *string   `json:"country_code3,omitempty" opensearch:"keyword"`
	CountryCode2   *string   `json:"country_code2,omitempty" opensearch:"keyword"`
	CountryName    *string   `json:"country_name,omitempty" opensearch:"keyword"`
	IP             *string   `json:"ip,omitempty" opensearch:"keyword"`
	Longitude      *float64  `json:"longitude,omitempty" opensearch:"float"`
	Timezone       *string   `json:"timezone,omitempty" opensearch:"keyword"`
	AreaCode       *uint64   `json:"area_code,omitempty" opensearch:"long"`
	CityName       *string   `json:"city_name,omitempty" opensearch:"keyword"`
	ContinentCode  *string   `json:"continent_code,omitempty" opensearch:"keyword"`
	PostalCode     *string   `json:"postal_code,omitempty" opensearch:"keyword"`
}

type Mitre struct {
	ID        *string `json:"id,omitempty"`
	Tactic    *string `json:"tactic,omitempty"`
	Technique *string `json:"technique,omitempty"`
}

type Rule struct {
	Description *string `json:"description,omitempty" opensearch:"keyword"`
	ID          *string `json:"id,omitempty" opensearch:"keyword"`
	Level       *uint64 `json:"level,omitempty" opensearch:"long"`
	Frequency   *uint64 `json:"frequency,omitempty" opensearch:"long"`
	Firedtimes  *uint64 `json:"firedtimes,omitempty" opensearch:"long"`
	Gdpr        *string `json:"gdpr,omitempty" opensearch:"keyword"`
	Gpg13       *string `json:"gpg13,omitempty" opensearch:"keyword"`
	Mail        *bool   `json:"mail,omitempty" opensearch:"keyword"`
	Groups      *string `json:"groups,omitempty" opensearch:"keyword"`
	Info        *string `json:"info,omitempty" opensearch:"keyword"`
	Cis         *string `json:"cis,omitempty" opensearch:"keyword"`
	PciDss      *string `json:"pci_dss,omitempty" opensearch:"keyword"`
	Hipaa       *string `json:"hipaa,omitempty" opensearch:"keyword"`
	Mitre       *Mitre  `json:"mitre,omitempty" opensearch:"object"`
	Cve         *string `json:"cve,omitempty" opensearch:"keyword"`
	Nist80053   *string `json:"nist_800_53,omitempty" opensearch:"keyword"`
}

type File struct {
	Inode *string `json:"inode,omitempty" opensearch:"keyword"`
	Mode  *string `json:"mode,omitempty" opensearch:"keyword"`
	Name  *string `json:"name,omitempty" opensearch:"keyword"`
}

type Directory struct {
	Name  *string `json:"name,omitempty" opensearch:"keyword"`
	Inode *string `json:"inode,omitempty" opensearch:"keyword"`
	Mode  *string `json:"mode,omitempty" opensearch:"keyword"`
}

type Execve struct {
	A0 *string `json:"a0,omitempty" opensearch:"keyword"`
	A1 *string `json:"a1,omitempty" opensearch:"keyword"`
	A2 *string `json:"a2,omitempty" opensearch:"keyword"`
	A3 *string `json:"a3,omitempty" opensearch:"keyword"`
}

type DataAudit struct {
	Dev          *string    `json:"dev,omitempty" opensearch:"keyword"`
	Ppid         *string    `json:"ppid,omitempty" opensearch:"keyword"`
	Subj         *string    `json:"subj,omitempty" opensearch:"keyword"`
	OldProm      *string    `json:"old_prom,omitempty" opensearch:"keyword"`
	Auid         *string    `json:"auid,omitempty" opensearch:"keyword"`
	Egid         *string    `json:"egid,omitempty" opensearch:"keyword"`
	Exe          *string    `json:"exe,omitempty" opensearch:"keyword"`
	Exit         *string    `json:"exit,omitempty" opensearch:"keyword"`
	Fsgid        *string    `json:"fsgid,omitempty" opensearch:"keyword"`
	OldEnforcing *string    `json:"old_enforcing,omitempty" opensearch:"keyword"`
	Key          *string    `json:"key,omitempty" opensearch:"keyword"`
	Res          *string    `json:"res,omitempty" opensearch:"keyword"`
	Session      *string    `json:"session,omitempty" opensearch:"keyword"`
	Success      *string    `json:"success,omitempty" opensearch:"keyword"`
	Acct         *string    `json:"acct,omitempty" opensearch:"keyword"`
	Command      *string    `json:"command,omitempty" opensearch:"keyword"`
	File         *File      `json:"file,omitempty" opensearch:"object"`
	Type         *string    `json:"type,omitempty" opensearch:"keyword"`
	Cwd          *string    `json:"cwd,omitempty" opensearch:"keyword"`
	Directory    *Directory `json:"directory,omitempty" opensearch:"object"`
	ID           *string    `json:"id,omitempty" opensearch:"keyword"`
	Syscall      *string    `json:"syscall,omitempty" opensearch:"keyword"`
	Arch         *string    `json:"arch,omitempty" opensearch:"keyword"`
	Euid         *string    `json:"euid,omitempty" opensearch:"keyword"`
	Fsuid        *string    `json:"fsuid,omitempty" opensearch:"keyword"`
	List         *string    `json:"list,omitempty" opensearch:"keyword"`
	Pid          *string    `json:"pid,omitempty" opensearch:"keyword"`
	Suid         *string    `json:"suid,omitempty" opensearch:"keyword"`
	Execve       *Execve    `json:"execve,omitempty" opensearch:"object"`
	Prom         *string    `json:"prom,omitempty" opensearch:"keyword"`
	Sgid         *string    `json:"sgid,omitempty" opensearch:"keyword"`
	Srcip        *string    `json:"srcip,omitempty" opensearch:"keyword"`
	Uid          *string    `json:"uid,omitempty" opensearch:"keyword"`
	Enforcing    *string    `json:"enforcing,omitempty" opensearch:"keyword"`
	Gid          *string    `json:"gid,omitempty" opensearch:"keyword"`
	Op           *string    `json:"op,omitempty" opensearch:"keyword"`
	Tty          *string    `json:"tty,omitempty" opensearch:"keyword"`
}

type Hardware struct {
	CpuName  *string  `json:"cpu_name,omitempty" opensearch:"keyword"`
	CpuCores *uint64  `json:"cpu_cores,omitempty" opensearch:"long"`
	CpuMhz   *float64 `json:"cpu_mhz,omitempty" opensearch:"float"`
	RamTotal *uint64  `json:"ram_total,omitempty" opensearch:"keyword"`
	RamFree  *uint64  `json:"ram_free,omitempty" opensearch:"long"`
	RamUsage *uint64  `json:"ram_usage,omitempty" opensearch:"long"`
	Serial   *string  `json:"serial,omitempty" opensearch:"keyword"`
}

type Compliance struct {
	Nist80053 *string `json:"nist_800_53,omitempty" opensearch:"keyword"`
	Cis       *string `json:"cis,omitempty" opensearch:"keyword"`
	CisCsc    *string `json:"cis_csc,omitempty" opensearch:"keyword"`
	PciDss    *string `json:"pci_dss,omitempty" opensearch:"keyword"`
	Hipaa     *string `json:"hipaa,omitempty" opensearch:"keyword"`
}

type ScaCheck struct {
	Title          *string     `json:"title,omitempty" opensearch:"keyword"`
	Description    *string     `json:"description,omitempty" opensearch:"keyword"`
	PreviousResult *string     `json:"previous_result,omitempty" opensearch:"keyword"`
	Reason         *string     `json:"reason,omitempty" opensearch:"keyword"`
	Status         *string     `json:"status,omitempty" opensearch:"keyword"`
	Rationale      *string     `json:"rationale,omitempty" opensearch:"keyword"`
	Compliance     *Compliance `json:"compliance,omitempty" opensearch:"keyword"`
	References     *string     `json:"references,omitempty" opensearch:"keyword"`
	Registry       *string     `json:"registry,omitempty" opensearch:"keyword"`
	Result         *string     `json:"result,omitempty" opensearch:"keyword"`
	ID             *string     `json:"id,omitempty" opensearch:"keyword"`
	File           *string     `json:"file,omitempty" opensearch:"keyword"`
	Directory      *string     `json:"directory,omitempty" opensearch:"keyword"`
	Remediation    *string     `json:"remediation,omitempty" opensearch:"keyword"`
	Process        *string     `json:"process,omitempty" opensearch:"keyword"`
}

type Sca struct {
	ScanId      *string   `json:"scan_id,omitempty" opensearch:"keyword"`
	Name        *string   `json:"name,omitempty" opensearch:"keyword"`
	Description *string   `json:"description,omitempty" opensearch:"keyword"`
	Passed      *int64    `json:"passed,omitempty" opensearch:"long"`
	Failed      *int64    `json:"failed,omitempty" opensearch:"long"`
	Score       *uint64   `json:"score,omitempty" opensearch:"long"`
	TotalChecks *string   `json:"total_checks,omitempty" opensearch:"keyword"`
	Type        *string   `json:"type,omitempty" opensearch:"keyword"`
	Policy      *string   `json:"policy,omitempty" opensearch:"keyword"`
	File        *string   `json:"file,omitempty" opensearch:"keyword"`
	Check       *ScaCheck `json:"sca_check,omitempty" opensearch:"object"`
	Invalid     *string   `json:"invalid,omitempty" opensearch:"keyword"`
	PolicyId    *string   `json:"policy_id,omitempty" opensearch:"keyword"`
}

type Oval struct {
	ID *string `json:"id,omitempty" opensearch:"keyword"`
}

type OscapCheck struct {
	Result      *string `json:"result,omitempty" opensearch:"keyword"`
	Identifiers *string `json:"identifiers,omitempty" opensearch:"keyword"`
	Oval        *Oval   `json:"oval,omitempty" opensearch:"object"`
	Rationale   *string `json:"rationale,omitempty" opensearch:"keyword"`
	Severity    *string `json:"severity,omitempty" opensearch:"keyword"`
	Title       *string `json:"title,omitempty" opensearch:"keyword"`
	Description *string `json:"description,omitempty" opensearch:"keyword"`
	ID          *string `json:"id,omitempty" opensearch:"keyword"`
	References  *string `json:"references,omitempty" opensearch:"keyword"`
}

type Benchmark struct {
	ID *string `json:"id,omitempty" opensearch:"keyword"`
}

type Profile struct {
	ID    *string `json:"id,omitempty" opensearch:"keyword"`
	Title *string `json:"title,omitempty" opensearch:"keyword"`
}

type Scan struct {
	Benchmark  *Benchmark `json:"benchmark,omitempty" opensearch:"object"`
	Content    *string    `json:"content,omitempty" opensearch:"keyword"`
	ID         *string    `json:"id,omitempty" opensearch:"keyword"`
	Profile    *Profile   `json:"profile,omitempty" opensearch:"object"`
	ReturnCode *uint64    `json:"return_code,omitempty" opensearch:"long"`
	Score      *float64   `json:"score,omitempty" opensearch:"float"`
}
type Oscap struct {
	Check *OscapCheck `json:"oscap_check,omitempty" opensearch:"object"`
	Scan  *Scan       `json:"scan,omitempty" opensearch:"object"`
}

type Package struct {
	Version      *string `json:"version,omitempty" opensearch:"keyword"`
	Source       *string `json:"source,omitempty" opensearch:"keyword"`
	Architecture *string `json:"architecture,omitempty" opensearch:"keyword"`
	Condition    *string `json:"condition,omitempty" opensearch:"keyword"`
	GeneratedCpe *string `json:"generated_cpe,omitempty" opensearch:"keyword"`
	Name         *string `json:"name,omitempty" opensearch:"keyword"`
}

type Cvss2Vector struct {
	AccessComplexity      *string `json:"access_complexity,omitempty" opensearch:"keyword"`
	Authentication        *string `json:"authentication,omitempty" opensearch:"keyword"`
	IntegrityImpact       *string `json:"integrity_impact,omitempty" opensearch:"keyword"`
	Scope                 *string `json:"scope,omitempty" opensearch:"keyword"`
	AttackVector          *string `json:"attack_vector,omitempty" opensearch:"keyword"`
	Availability          *string `json:"availability,omitempty" opensearch:"keyword"`
	ConfidentialityImpact *string `json:"confidentiality_impact,omitempty" opensearch:"keyword"`
	PrivilegesRequired    *string `json:"privileges_required,omitempty" opensearch:"keyword"`
	UserInteraction       *string `json:"user_interaction,omitempty" opensearch:"keyword"`
}

type Cvss2 struct {
	BaseScore           *string      `json:"base_score,omitempty" opensearch:"keyword"`
	ExploitabilityScore *string      `json:"exploitability_score,omitempty" opensearch:"keyword"`
	ImpactScore         *string      `json:"impact_score,omitempty" opensearch:"keyword"`
	Vector              *Cvss2Vector `json:"cvss2_vector,omitempty" opensearch:"object"`
}

type Cvss3Vector struct {
	AttackVector          *string `json:"attack_vector,omitempty" opensearch:"keyword"`
	Scope                 *string `json:"scope,omitempty" opensearch:"keyword"`
	ConfidentialityImpact *string `json:"confidentiality_impact,omitempty" opensearch:"keyword"`
	IntegrityImpact       *string `json:"integrity_impact,omitempty" opensearch:"keyword"`
	PrivilegesRequired    *string `json:"privileges_required,omitempty" opensearch:"keyword"`
	UserInteraction       *string `json:"user_interaction,omitempty" opensearch:"keyword"`
	AccessComplexity      *string `json:"access_complexity,omitempty" opensearch:"keyword"`
	Authentication        *string `json:"authentication,omitempty" opensearch:"keyword"`
	Availability          *string `json:"availability,omitempty" opensearch:"keyword"`
}

type Cvss3 struct {
	ExploitabilityScore *string      `json:"exploitability_score,omitempty" opensearch:"keyword"`
	ImpactScore         *string      `json:"impact_score,omitempty" opensearch:"keyword"`
	Vector              *Cvss3Vector `json:"cvss3_vector,omitempty" opensearch:"object"`
	BaseScore           *string      `json:"base_score,omitempty" opensearch:"keyword"`
}

type Cvss struct {
	Cvss2 *Cvss2 `json:"cvss2,omitempty" opensearch:"object"`
	Cvss3 *Cvss3 `json:"cvss3,omitempty" opensearch:"object"`
}

type Vulnerability struct {
	Cve          *string  `json:"cve,omitempty" opensearch:"keyword"`
	Package      *Package `json:"package,omitempty" opensearch:"object"`
	Published    *int64   `json:"published,omitempty" opensearch:"long"`
	Rationale    *string  `json:"rationale,omitempty" opensearch:"keyword"`
	Severity     *string  `json:"severity,omitempty" opensearch:"keyword"`
	Assigner     *string  `json:"assigner,omitempty" opensearch:"keyword"`
	CveVersion   *string  `json:"cve_version,omitempty" opensearch:"keyword"`
	Cvss         *Cvss    `json:"cvss,omitempty" opensearch:"object"`
	CweReference *string  `json:"cwe_reference,omitempty" opensearch:"keyword"`
	Updated      *int64   `json:"updated,omitempty" opensearch:"long"`
	Title        *string  `json:"title,omitempty" opensearch:"keyword"`
}

type NetworkInterfaces struct {
	PrivateIpAddress *string `json:"privateIpAddress,omitempty" opensearch:"keyword"`
	PublicIp         *string `json:"publicIp,omitempty" opensearch:"keyword"`
}

type ResourceInstanceDetails struct {
	LaunchTime        *int64             `json:"launchTime,omitempty" opensearch:"long"`
	NetworkInterfaces *NetworkInterfaces `json:"networkInterfaces,omitempty" opensearch:"keyword"`
}

type ActionNetworkConnectionActionRemoteIpDetails struct {
	IpAddressV4 *string   `json:"ipAddressV4,omitempty" opensearch:"keyword"`
	GeoLocation *GeoPoint `json:"geoLocation,omitempty" opensearch:"geo_point"`
}

type Service struct {
	Count                                        *uint64                                       `json:"count,omitempty" opensearch:"long"`
	ActionNetworkConnectionActionRemoteIpDetails *ActionNetworkConnectionActionRemoteIpDetails `json:"action.networkConnectionAction.remoteIpDetails,omitempty" opensearch:"object"`
	EventFirstSeen                               *int64                                        `json:"eventFirstSeen,omitempty" opensearch:"long"`
	EventLastSeen                                *int64                                        `json:"eventLastSeen,omitempty" opensearch:"long"`
}

type Aws struct {
	Bytes                   *uint64                  `json:"bytes,omitempty" opensearch:"long"`
	Srcaddr                 *string                  `json:"srcaddr,omitempty" opensearch:"keyword"`
	End                     *int64                   `json:"end,omitempty" opensearch:"long"`
	Start                   *int64                   `json:"start,omitempty" opensearch:"long"`
	SourceIpAddress         *string                  `json:"source_ip_address,omitempty" opensearch:"keyword"`
	ResourceInstanceDetails *ResourceInstanceDetails `json:"resource.instanceDetails,omitempty" opensearch:"object"`
	Dstaddr                 *string                  `json:"dstaddr,omitempty" opensearch:"keyword"`
	Service                 *Service                 `json:"service,omitempty" opensearch:"object"`
	CreatedAt               *int64                   `json:"createdAt,omitempty" opensearch:"long"`
	UpdatedAt               *int64                   `json:"updatedAt,omitempty" opensearch:"long"`
}

type Source struct {
	AlertId *string `json:"alert_id,omitempty" opensearch:"keyword"`
	File    *string `json:"file,omitempty" opensearch:"keyword"`
	Md5     *string `json:"md5,omitempty" opensearch:"keyword"`
	Sha1    *string `json:"sha1,omitempty" opensearch:"keyword"`
}

type Virustotal struct {
	ScanDate    *string `json:"scan_date,omitempty" opensearch:"keyword"`
	Sha1        *string `json:"sha1,omitempty" opensearch:"keyword"`
	Source      *Source `json:"source,omitempty" opensearch:"object"`
	Total       *string `json:"total,omitempty" opensearch:"keyword"`
	Description *string `json:"description,omitempty" opensearch:"keyword"`
	Error       *string `json:"error,omitempty" opensearch:"keyword"`
	Malicious   *string `json:"malicious,omitempty" opensearch:"keyword"`
	Positives   *string `json:"positives,omitempty" opensearch:"keyword"`
	Found       *string `json:"found,omitempty" opensearch:"keyword"`
	Permalink   *string `json:"permalink,omitempty" opensearch:"keyword"`
}

type DataProcess struct {
	Priority  *uint64 `json:"priority,omitempty" opensearch:"long"`
	Ppid      *uint64 `json:"ppid,omitempty" opensearch:"long"`
	Utime     *uint64 `json:"utime,omitempty" opensearch:"long"`
	Size      *uint64 `json:"size,omitempty" opensearch:"long"`
	VmSize    *uint64 `json:"vm_size,omitempty" opensearch:"long"`
	Session   *uint64 `json:"session,omitempty" opensearch:"long"`
	State     *string `json:"state,omitempty" opensearch:"keyword"`
	Stime     *uint64 `json:"stime,omitempty" opensearch:"long"`
	Pid       *uint64 `json:"pid,omitempty" opensearch:"long"`
	Fgroup    *string `json:"fgroup,omitempty" opensearch:"keyword"`
	Tgid      *uint64 `json:"tgid,omitempty" opensearch:"long"`
	Rgroup    *string `json:"rgroup,omitempty" opensearch:"keyword"`
	Share     *uint64 `json:"share,omitempty" opensearch:"long"`
	Nlwp      *uint64 `json:"nlwp,omitempty" opensearch:"long"`
	Tty       *uint64 `json:"tty,omitempty" opensearch:"long"`
	Args      *string `json:"args,omitempty" opensearch:"keyword"`
	Euser     *string `json:"euser,omitempty" opensearch:"keyword"`
	Suser     *string `json:"suser,omitempty" opensearch:"keyword"`
	StartTime *uint64 `json:"start_time,omitempty" opensearch:"long"`
	Pgrp      *uint64 `json:"pgrp,omitempty" opensearch:"long"`
	Name      *string `json:"name,omitempty" opensearch:"keyword"`
	Cmd       *string `json:"cmd,omitempty" opensearch:"keyword"`
	Processor *uint64 `json:"processor,omitempty" opensearch:"long"`
	Sgroup    *string `json:"sgroup,omitempty" opensearch:"keyword"`
	Resident  *uint64 `json:"resident,omitempty" opensearch:"long"`
	Nice      *uint64 `json:"nice,omitempty" opensearch:"long"`
	Ruser     *string `json:"ruser,omitempty" opensearch:"keyword"`
	Egroup    *string `json:"egroup,omitempty" opensearch:"keyword"`
}

type Ipv6 struct {
	Netmask   *string `json:"netmask,omitempty" opensearch:"keyword"`
	Broadcast *string `json:"broadcast,omitempty" opensearch:"keyword"`
	Metric    *uint64 `json:"metric,omitempty" opensearch:"long"`
	Gateway   *string `json:"gateway,omitempty" opensearch:"keyword"`
	Dhcp      *string `json:"dhcp,omitempty" opensearch:"keyword"`
	Address   *string `json:"address,omitempty" opensearch:"keyword"`
}

type Ipv4 struct {
	Address   *string `json:"address,omitempty" opensearch:"keyword"`
	Netmask   *string `json:"netmask,omitempty" opensearch:"keyword"`
	Broadcast *string `json:"broadcast,omitempty" opensearch:"keyword"`
	Metric    *uint64 `json:"metric,omitempty" opensearch:"long"`
	Gateway   *string `json:"gateway,omitempty" opensearch:"keyword"`
	Dhcp      *string `json:"dhcp,omitempty" opensearch:"keyword"`
}

type Iface struct {
	Type      *string `json:"type,omitempty" opensearch:"keyword"`
	RxBytes   *uint64 `json:"rx_bytes,omitempty" opensearch:"long"`
	TxPackets *uint64 `json:"tx_packets,omitempty" opensearch:"long"`
	Ipv6      *Ipv6   `json:"ipv6,omitempty" opensearch:"object"`
	Name      *string `json:"name,omitempty" opensearch:"keyword"`
	State     *string `json:"state,omitempty" opensearch:"keyword"`
	Mtu       *uint64 `json:"mtu,omitempty" opensearch:"long"`
	RxPackets *uint64 `json:"rx_packets,omitempty" opensearch:"long"`
	Ipv4      *Ipv4   `json:"ipv4,omitempty" opensearch:"object"`
	Adapter   *string `json:"adapter,omitempty" opensearch:"keyword"`
	TxBytes   *uint64 `json:"tx_bytes,omitempty" opensearch:"long"`
	TxErrors  *uint64 `json:"tx_errors,omitempty" opensearch:"long"`
	RxErrors  *uint64 `json:"rx_errors,omitempty" opensearch:"long"`
	TxDropped *uint64 `json:"tx_dropped,omitempty" opensearch:"long"`
	Mac       *string `json:"mac,omitempty" opensearch:"keyword"`
	RxDropped *uint64 `json:"rx_dropped,omitempty" opensearch:"long"`
}

type Netinfo struct {
	Iface *Iface `json:"iface,omitempty" opensearch:"object"`
}

type Os struct {
	Version        *string `json:"version,omitempty" opensearch:"keyword"`
	Major          *string `json:"major,omitempty" opensearch:"keyword"`
	Minor          *string `json:"minor,omitempty" opensearch:"keyword"`
	Platform       *string `json:"platform,omitempty" opensearch:"keyword"`
	Release        *string `json:"release,omitempty" opensearch:"keyword"`
	ReleaseVersion *string `json:"release_version,omitempty" opensearch:"keyword"`
	Architecture   *string `json:"architecture,omitempty" opensearch:"keyword"`
	Name           *string `json:"name,omitempty" opensearch:"keyword"`
	Build          *string `json:"build,omitempty" opensearch:"keyword"`
	Sysname        *string `json:"sysname,omitempty" opensearch:"keyword"`
	Hostname       *string `json:"hostname,omitempty" opensearch:"keyword"`
	Codename       *string `json:"codename,omitempty" opensearch:"keyword"`
}

type Port struct {
	RemoteIp   *string `json:"remote_ip,omitempty" opensearch:"keyword"`
	TxQueue    *uint64 `json:"tx_queue,omitempty" opensearch:"long"`
	Inode      *uint64 `json:"inode,omitempty" opensearch:"long"`
	State      *string `json:"state,omitempty" opensearch:"keyword"`
	Pid        *uint64 `json:"pid,omitempty" opensearch:"long"`
	Process    *string `json:"process,omitempty" opensearch:"keyword"`
	LocalPort  *uint64 `json:"local_port,omitempty" opensearch:"long"`
	LocalIp    *string `json:"local_ip,omitempty" opensearch:"keyword"`
	RemotePort *uint64 `json:"remote_port,omitempty"`
	RxQueue    *uint64 `json:"rx_queue,omitempty" opensearch:"long"`
	Protocol   *string `json:"protocol,omitempty" opensearch:"keyword"`
}

type Program struct {
	Version      *string `json:"version,omitempty" opensearch:"keyword"`
	Architecture *string `json:"architecture,omitempty" opensearch:"keyword"`
	Multiarch    *string `json:"multiarch,omitempty" opensearch:"keyword"`
	Source       *string `json:"source,omitempty" opensearch:"keyword"`
	Description  *string `json:"description,omitempty" opensearch:"keyword"`
	Priority     *string `json:"priority,omitempty" opensearch:"keyword"`
	InstallTime  *string `json:"install_time,omitempty" opensearch:"keyword"`
	Section      *string `json:"section,omitempty" opensearch:"keyword"`
	Size         *uint64 `json:"size,omitempty" opensearch:"long"`
	Vendor       *string `json:"vendor,omitempty" opensearch:"keyword"`
	Location     *string `json:"location,omitempty" opensearch:"keyword"`
	Format       *string `json:"format,omitempty" opensearch:"keyword"`
	Name         *string `json:"name,omitempty" opensearch:"keyword"`
}

type Data struct {
	Audit         *DataAudit     `json:"data_audit,omitempty" opensearch:"object"`
	Status        *string        `json:"status,omitempty" opensearch:"keyword"`
	Uid           *string        `json:"uid,omitempty" opensearch:"keyword"`
	Dstip         *string        `json:"dstip,omitempty" opensearch:"keyword"`
	Hardware      *Hardware      `json:"hardware,omitempty" opensearch:"object"`
	Sca           *Sca           `json:"sca,omitempty" opensearch:"object"`
	Command       *string        `json:"command,omitempty" opensearch:"keyword"`
	Timestamp     *int64         `json:"timestamp,omitempty" opensearch:"long"`
	Title         *string        `json:"title,omitempty" opensearch:"keyword"`
	Srcuser       *string        `json:"srcuser,omitempty" opensearch:"keyword"`
	ID            *string        `json:"id,omitempty" opensearch:"keyword"`
	SystemName    *string        `json:"system_name,omitempty" opensearch:"keyword"`
	Oscap         *Oscap         `json:"oscap,omitempty" opensearch:"object"`
	Dstport       *string        `json:"dstport,omitempty" opensearch:"keyword"`
	Dstuser       *string        `json:"dstuser,omitempty" opensearch:"keyword"`
	Vulnerability *Vulnerability `json:"vulnerability,omitempty" opensearch:"object"`
	Aws           *Aws           `json:"aws,omitempty" opensearch:"object"`
	Action        *string        `json:"action,omitempty" opensearch:"keyword"`
	Virustotal    *Virustotal    `json:"virustotal,omitempty" opensearch:"object"`
	Srcip         *string        `json:"srcip,omitempty" opensearch:"keyword"`
	ExtraData     *string        `json:"extra_data,omitempty" opensearch:"keyword"`
	Integration   *string        `json:"integration,omitempty" opensearch:"keyword"`
	Process       *DataProcess   `json:"data_process,omitempty" opensearch:"object"`
	Data          *string        `json:"data,omitempty" opensearch:"keyword"`
	URL           *string        `json:"url,omitempty" opensearch:"keyword"`
	Type          *string        `json:"type,omitempty" opensearch:"keyword"`
	Netinfo       *Netinfo       `json:"netinfo,omitempty" opensearch:"object"`
	Os            *Os            `json:"os,omitempty" opensearch:"object"`
	Port          *Port          `json:"port,omitempty" opensearch:"object"`
	Program       *Program       `json:"program,omitempty" opensearch:"object"`
	Protocol      *string        `json:"protocol,omitempty" opensearch:"keyword"`
	Srcport       *string        `json:"srcport,omitempty" opensearch:"keyword"`
}

type Alert struct {
	Manager        *Manager     `json:"manager,omitempty" opensearch:"object"`
	Cluster        *Cluster     `json:"cluster,omitempty" opensearch:"object"`
	Syscheck       *Syscheck    `json:"syscheck,omitempty" opensearch:"object"`
	Location       *string      `json:"location,omitempty" opensearch:"keyword"`
	Decoder        *Decoder     `json:"decoder,omitempty" opensearch:"object"`
	Offset         *string      `json:"offset,omitempty" opensearch:"keyword"`
	Command        *string      `json:"command,omitempty" opensearch:"keyword"`
	Type           *string      `json:"type,omitempty" opensearch:"keyword"`
	Timestamp_     *int64       `json:"@timestamp,omitempty" opensearch:"long"`
	Agent          *Agent       `json:"agent,omitempty" opensearch:"object"`
	FullLog        *string      `json:"full_log,omitempty" opensearch:"keyword"`
	PreviousLog    *string      `json:"previous_log,omitempty" opensearch:"keyword"`
	Message        *string      `json:"message,omitempty" opensearch:"keyword"`
	Input          *Input       `json:"input,omitempty" opensearch:"object"`
	Timestamp      *int64       `json:"timestamp,omitempty" opensearch:"long"`
	Version_       *string      `json:"@version,omitempty" opensearch:"keyword"`
	Host           *string      `json:"host,omitempty" opensearch:"keyword"`
	Predecoder     *Predecoder  `json:"predecoder,omitempty" opensearch:"object"`
	ID             *string      `json:"id,omitempty" opensearch:"keyword"`
	PreviousOutput *string      `json:"previous_output,omitempty" opensearch:"keyword"`
	GeoLocation    *GeoLocation `json:"GeoLocation,omitempty" opensearch:"geo_point"`
	Rule           *Rule        `json:"rule,omitempty" opensearch:"object"`
	Data           *Data        `json:"data,omitempty" opensearch:"object"`
	ProgramName    *string      `json:"program_name,omitempty" opensearch:"keyword"`
	Title          *string      `json:"title,omitempty" opensearch:"keyword"`
}
