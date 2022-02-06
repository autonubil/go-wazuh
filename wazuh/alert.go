package wazuh

type GeoPoint struct {
	Lat  float64 `json:"lat,omitempty"`
	Long float64 `json:"long,omitempty"`
}

type Manager struct {
	Name *string `json:"name,omitempty"`
}

type Cluster struct {
	Name *string `json:"name,omitempty"`
	Node *string `json:"node,omitempty"`
}

type EffectiveUser struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Group struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type LoginUser struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SyscheckAuditProcess struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Ppid *string `json:"ppid,omitempty"`
}

type User struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type SyscheckAudit struct {
	EffectiveUser *EffectiveUser        `json:"effective_user,omitempty"`
	Group         *Group                `json:"group,omitempty"`
	LoginUser     *LoginUser            `json:"login_user,omitempty"`
	Process       *SyscheckAuditProcess `json:"syscheck_audit_process,omitempty"`
	User          *User                 `json:"user,omitempty"`
}

type Syscheck struct {
	UidAfter     *string        `json:"uid_after,omitempty"`
	Md5Before    *string        `json:"md5_before,omitempty"`
	MtimeAfter   *int64         `json:"mtime_after,omitempty"`
	GidAfter     *string        `json:"gid_after,omitempty"`
	Diff         *string        `json:"diff,omitempty"`
	Path         *string        `json:"path,omitempty"`
	PermAfter    *string        `json:"perm_after,omitempty"`
	InodeAfter   *string        `json:"inode_after,omitempty"`
	InodeBefore  *string        `json:"inode_before,omitempty"`
	MtimeBefore  *int64         `json:"mtime_before,omitempty"`
	UnameBefore  *string        `json:"uname_before,omitempty"`
	Sha256After  *string        `json:"sha256_after,omitempty"`
	HardLinks    *string        `json:"hard_links,omitempty"`
	UidBefore    *string        `json:"uid_before,omitempty"`
	Event        *string        `json:"event,omitempty"`
	Sha256Before *string        `json:"sha256_before,omitempty"`
	Sha1After    *string        `json:"sha1_after,omitempty"`
	GnameBefore  *string        `json:"gname_before,omitempty"`
	Tags         *string        `json:"tags,omitempty"`
	Sha1Before   *string        `json:"sha1_before,omitempty"`
	PermBefore   *string        `json:"perm_before,omitempty"`
	GnameAfter   *string        `json:"gname_after,omitempty"`
	UnameAfter   *string        `json:"uname_after,omitempty"`
	SizeAfter    *uint64        `json:"size_after,omitempty"`
	Mode         *string        `json:"mode,omitempty"`
	GidBefore    *string        `json:"gid_before,omitempty"`
	Md5After     *string        `json:"md5_after,omitempty"`
	SizeBefore   *uint64        `json:"size_before,omitempty"`
	Audit        *SyscheckAudit `json:"syscheck_audit,omitempty"`
}

type Decoder struct {
	Accumulate *uint64 `json:"accumulate,omitempty"`
	Parent     *string `json:"parent,omitempty"`
	Name       *string `json:"name,omitempty"`
	Ftscomment *string `json:"ftscomment,omitempty"`
	Fts        *uint64 `json:"fts,omitempty"`
}

type Agent struct {
	IP   *string `json:"ip,omitempty"`
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Input struct {
	Type *string `json:"type,omitempty"`
}

type Predecoder struct {
	ProgramName *string `json:"program_name,omitempty"`
	Timestamp   *string `json:"timestamp,omitempty"`
	Hostname    *string `json:"hostname,omitempty"`
}

type GeoLocation struct {
	RegionName     *string   `json:"region_name,omitempty"`
	Latitude       *float64  `json:"latitude,omitempty"`
	RealRegionName *string   `json:"real_region_name,omitempty"`
	DmaCode        *uint64   `json:"dma_code,omitempty"`
	Location       *GeoPoint `json:"location,omitempty"`
	Coordinates    *float64  `json:"coordinates,omitempty"`
	CountryCode3   *string   `json:"country_code3,omitempty"`
	CountryCode2   *string   `json:"country_code2,omitempty"`
	CountryName    *string   `json:"country_name,omitempty"`
	IP             *string   `json:"ip,omitempty"`
	Longitude      *float64  `json:"longitude,omitempty"`
	Timezone       *string   `json:"timezone,omitempty"`
	AreaCode       *uint64   `json:"area_code,omitempty"`
	CityName       *string   `json:"city_name,omitempty"`
	ContinentCode  *string   `json:"continent_code,omitempty"`
	PostalCode     *string   `json:"postal_code,omitempty"`
}

type Mitre struct {
	ID        *string `json:"id,omitempty"`
	Tactic    *string `json:"tactic,omitempty"`
	Technique *string `json:"technique,omitempty"`
}

type Rule struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Level       *uint64 `json:"level,omitempty"`
	Frequency   *uint64 `json:"frequency,omitempty"`
	Firedtimes  *uint64 `json:"firedtimes,omitempty"`
	Gdpr        *string `json:"gdpr,omitempty"`
	Gpg13       *string `json:"gpg13,omitempty"`
	Mail        *bool   `json:"mail,omitempty"`
	Groups      *string `json:"groups,omitempty"`
	Info        *string `json:"info,omitempty"`
	Cis         *string `json:"cis,omitempty"`
	PciDss      *string `json:"pci_dss,omitempty"`
	Hipaa       *string `json:"hipaa,omitempty"`
	Mitre       *Mitre  `json:"mitre,omitempty"`
	Cve         *string `json:"cve,omitempty"`
	Nist80053   *string `json:"nist_800_53,omitempty"`
}

type File struct {
	Inode *string `json:"inode,omitempty"`
	Mode  *string `json:"mode,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type Directory struct {
	Name  *string `json:"name,omitempty"`
	Inode *string `json:"inode,omitempty"`
	Mode  *string `json:"mode,omitempty"`
}

type Execve struct {
	A0 *string `json:"a0,omitempty"`
	A1 *string `json:"a1,omitempty"`
	A2 *string `json:"a2,omitempty"`
	A3 *string `json:"a3,omitempty"`
}

type DataAudit struct {
	Dev          *string    `json:"dev,omitempty"`
	Ppid         *string    `json:"ppid,omitempty"`
	Subj         *string    `json:"subj,omitempty"`
	OldProm      *string    `json:"old_prom,omitempty"`
	Auid         *string    `json:"auid,omitempty"`
	Egid         *string    `json:"egid,omitempty"`
	Exe          *string    `json:"exe,omitempty"`
	Exit         *string    `json:"exit,omitempty"`
	Fsgid        *string    `json:"fsgid,omitempty"`
	OldEnforcing *string    `json:"old_enforcing,omitempty"`
	Key          *string    `json:"key,omitempty"`
	Res          *string    `json:"res,omitempty"`
	Session      *string    `json:"session,omitempty"`
	Success      *string    `json:"success,omitempty"`
	Acct         *string    `json:"acct,omitempty"`
	Command      *string    `json:"command,omitempty"`
	File         *File      `json:"file,omitempty"`
	Type         *string    `json:"type,omitempty"`
	Cwd          *string    `json:"cwd,omitempty"`
	Directory    *Directory `json:"directory,omitempty"`
	ID           *string    `json:"id,omitempty"`
	Syscall      *string    `json:"syscall,omitempty"`
	Arch         *string    `json:"arch,omitempty"`
	Euid         *string    `json:"euid,omitempty"`
	Fsuid        *string    `json:"fsuid,omitempty"`
	List         *string    `json:"list,omitempty"`
	Pid          *string    `json:"pid,omitempty"`
	Suid         *string    `json:"suid,omitempty"`
	Execve       *Execve    `json:"execve,omitempty"`
	Prom         *string    `json:"prom,omitempty"`
	Sgid         *string    `json:"sgid,omitempty"`
	Srcip        *string    `json:"srcip,omitempty"`
	Uid          *string    `json:"uid,omitempty"`
	Enforcing    *string    `json:"enforcing,omitempty"`
	Gid          *string    `json:"gid,omitempty"`
	Op           *string    `json:"op,omitempty"`
	Tty          *string    `json:"tty,omitempty"`
}

type Hardware struct {
	CpuName  *string  `json:"cpu_name,omitempty"`
	CpuCores *uint64  `json:"cpu_cores,omitempty"`
	CpuMhz   *float64 `json:"cpu_mhz,omitempty"`
	RamTotal *uint64  `json:"ram_total,omitempty"`
	RamFree  *uint64  `json:"ram_free,omitempty"`
	RamUsage *uint64  `json:"ram_usage,omitempty"`
	Serial   *string  `json:"serial,omitempty"`
}

type Compliance struct {
	Nist80053 *string `json:"nist_800_53,omitempty"`
	Cis       *string `json:"cis,omitempty"`
	CisCsc    *string `json:"cis_csc,omitempty"`
	PciDss    *string `json:"pci_dss,omitempty"`
	Hipaa     *string `json:"hipaa,omitempty"`
}

type ScaCheck struct {
	Title          *string     `json:"title,omitempty"`
	Description    *string     `json:"description,omitempty"`
	PreviousResult *string     `json:"previous_result,omitempty"`
	Reason         *string     `json:"reason,omitempty"`
	Status         *string     `json:"status,omitempty"`
	Rationale      *string     `json:"rationale,omitempty"`
	Compliance     *Compliance `json:"compliance,omitempty"`
	References     *string     `json:"references,omitempty"`
	Registry       *string     `json:"registry,omitempty"`
	Result         *string     `json:"result,omitempty"`
	ID             *string     `json:"id,omitempty"`
	File           *string     `json:"file,omitempty"`
	Directory      *string     `json:"directory,omitempty"`
	Remediation    *string     `json:"remediation,omitempty"`
	Process        *string     `json:"process,omitempty"`
}

type Sca struct {
	ScanId      *string   `json:"scan_id,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Passed      *int64    `json:"passed,omitempty"`
	Failed      *int64    `json:"failed,omitempty"`
	Score       *uint64   `json:"score,omitempty"`
	TotalChecks *string   `json:"total_checks,omitempty"`
	Type        *string   `json:"type,omitempty"`
	Policy      *string   `json:"policy,omitempty"`
	File        *string   `json:"file,omitempty"`
	Check       *ScaCheck `json:"sca_check,omitempty"`
	Invalid     *string   `json:"invalid,omitempty"`
	PolicyId    *string   `json:"policy_id,omitempty"`
}

type Oval struct {
	ID *string `json:"id,omitempty"`
}

type OscapCheck struct {
	Result      *string `json:"result,omitempty"`
	Identifiers *string `json:"identifiers,omitempty"`
	Oval        *Oval   `json:"oval,omitempty"`
	Rationale   *string `json:"rationale,omitempty"`
	Severity    *string `json:"severity,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	References  *string `json:"references,omitempty"`
}

type Benchmark struct {
	ID *string `json:"id,omitempty"`
}

type Profile struct {
	ID    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
}

type Scan struct {
	Benchmark  *Benchmark `json:"benchmark,omitempty"`
	Content    *string    `json:"content,omitempty"`
	ID         *string    `json:"id,omitempty"`
	Profile    *Profile   `json:"profile,omitempty"`
	ReturnCode *uint64    `json:"return_code,omitempty"`
	Score      *float64   `json:"score,omitempty"`
}
type Oscap struct {
	Check *OscapCheck `json:"oscap_check,omitempty"`
	Scan  *Scan       `json:"scan,omitempty"`
}

type Package struct {
	Version      *string `json:"version,omitempty"`
	Source       *string `json:"source,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Condition    *string `json:"condition,omitempty"`
	GeneratedCpe *string `json:"generated_cpe,omitempty"`
	Name         *string `json:"name,omitempty"`
}

type Cvss2Vector struct {
	AccessComplexity      *string `json:"access_complexity,omitempty"`
	Authentication        *string `json:"authentication,omitempty"`
	IntegrityImpact       *string `json:"integrity_impact,omitempty"`
	Scope                 *string `json:"scope,omitempty"`
	AttackVector          *string `json:"attack_vector,omitempty"`
	Availability          *string `json:"availability,omitempty"`
	ConfidentialityImpact *string `json:"confidentiality_impact,omitempty"`
	PrivilegesRequired    *string `json:"privileges_required,omitempty"`
	UserInteraction       *string `json:"user_interaction,omitempty"`
}

type Cvss2 struct {
	BaseScore           *string      `json:"base_score,omitempty"`
	ExploitabilityScore *string      `json:"exploitability_score,omitempty"`
	ImpactScore         *string      `json:"impact_score,omitempty"`
	Vector              *Cvss2Vector `json:"cvss2_vector,omitempty"`
}

type Cvss3Vector struct {
	AttackVector          *string `json:"attack_vector,omitempty"`
	Scope                 *string `json:"scope,omitempty"`
	ConfidentialityImpact *string `json:"confidentiality_impact,omitempty"`
	IntegrityImpact       *string `json:"integrity_impact,omitempty"`
	PrivilegesRequired    *string `json:"privileges_required,omitempty"`
	UserInteraction       *string `json:"user_interaction,omitempty"`
	AccessComplexity      *string `json:"access_complexity,omitempty"`
	Authentication        *string `json:"authentication,omitempty"`
	Availability          *string `json:"availability,omitempty"`
}

type Cvss3 struct {
	ExploitabilityScore *string      `json:"exploitability_score,omitempty"`
	ImpactScore         *string      `json:"impact_score,omitempty"`
	Vector              *Cvss3Vector `json:"cvss3_vector,omitempty"`
	BaseScore           *string      `json:"base_score,omitempty"`
}

type Cvss struct {
	Cvss2 *Cvss2 `json:"cvss2,omitempty"`
	Cvss3 *Cvss3 `json:"cvss3,omitempty"`
}

type Vulnerability struct {
	Cve          *string  `json:"cve,omitempty"`
	Package      *Package `json:"package,omitempty"`
	Published    *int64   `json:"published,omitempty"`
	Rationale    *string  `json:"rationale,omitempty"`
	Severity     *string  `json:"severity,omitempty"`
	Assigner     *string  `json:"assigner,omitempty"`
	CveVersion   *string  `json:"cve_version,omitempty"`
	Cvss         *Cvss    `json:"cvss,omitempty"`
	CweReference *string  `json:"cwe_reference,omitempty"`
	Updated      *int64   `json:"updated,omitempty"`
	Title        *string  `json:"title,omitempty"`
}

type NetworkInterfaces struct {
	PrivateIpAddress *string `json:"privateIpAddress,omitempty"`
	PublicIp         *string `json:"publicIp,omitempty"`
}

type ResourceInstanceDetails struct {
	LaunchTime        *int64             `json:"launchTime,omitempty"`
	NetworkInterfaces *NetworkInterfaces `json:"networkInterfaces,omitempty"`
}

type ActionNetworkConnectionActionRemoteIpDetails struct {
	IpAddressV4 *string   `json:"ipAddressV4,omitempty"`
	GeoLocation *GeoPoint `json:"geoLocation,omitempty"`
}

type Service struct {
	Count                                        *uint64                                       `json:"count,omitempty"`
	ActionNetworkConnectionActionRemoteIpDetails *ActionNetworkConnectionActionRemoteIpDetails `json:"action.networkConnectionAction.remoteIpDetails,omitempty"`
	EventFirstSeen                               *int64                                        `json:"eventFirstSeen,omitempty"`
	EventLastSeen                                *int64                                        `json:"eventLastSeen,omitempty"`
}

type Aws struct {
	Bytes                   *uint64                  `json:"bytes,omitempty"`
	Srcaddr                 *string                  `json:"srcaddr,omitempty"`
	End                     *int64                   `json:"end,omitempty"`
	Start                   *int64                   `json:"start,omitempty"`
	SourceIpAddress         *string                  `json:"source_ip_address,omitempty"`
	ResourceInstanceDetails *ResourceInstanceDetails `json:"resource.instanceDetails,omitempty"`
	Dstaddr                 *string                  `json:"dstaddr,omitempty"`
	Service                 *Service                 `json:"service,omitempty"`
	CreatedAt               *int64                   `json:"createdAt,omitempty"`
	UpdatedAt               *int64                   `json:"updatedAt,omitempty"`
}

type Source struct {
	AlertId *string `json:"alert_id,omitempty"`
	File    *string `json:"file,omitempty"`
	Md5     *string `json:"md5,omitempty"`
	Sha1    *string `json:"sha1,omitempty"`
}

type Virustotal struct {
	ScanDate    *string `json:"scan_date,omitempty"`
	Sha1        *string `json:"sha1,omitempty"`
	Source      *Source `json:"source,omitempty"`
	Total       *string `json:"total,omitempty"`
	Description *string `json:"description,omitempty"`
	Error       *string `json:"error,omitempty"`
	Malicious   *string `json:"malicious,omitempty"`
	Positives   *string `json:"positives,omitempty"`
	Found       *string `json:"found,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
}

type DataProcess struct {
	Priority  *uint64 `json:"priority,omitempty"`
	Ppid      *uint64 `json:"ppid,omitempty"`
	Utime     *uint64 `json:"utime,omitempty"`
	Size      *uint64 `json:"size,omitempty"`
	VmSize    *uint64 `json:"vm_size,omitempty"`
	Session   *uint64 `json:"session,omitempty"`
	State     *string `json:"state,omitempty"`
	Stime     *uint64 `json:"stime,omitempty"`
	Pid       *uint64 `json:"pid,omitempty"`
	Fgroup    *string `json:"fgroup,omitempty"`
	Tgid      *uint64 `json:"tgid,omitempty"`
	Rgroup    *string `json:"rgroup,omitempty"`
	Share     *uint64 `json:"share,omitempty"`
	Nlwp      *uint64 `json:"nlwp,omitempty"`
	Tty       *uint64 `json:"tty,omitempty"`
	Args      *string `json:"args,omitempty"`
	Euser     *string `json:"euser,omitempty"`
	Suser     *string `json:"suser,omitempty"`
	StartTime *uint64 `json:"start_time,omitempty"`
	Pgrp      *uint64 `json:"pgrp,omitempty"`
	Name      *string `json:"name,omitempty"`
	Cmd       *string `json:"cmd,omitempty"`
	Processor *uint64 `json:"processor,omitempty"`
	Sgroup    *string `json:"sgroup,omitempty"`
	Resident  *uint64 `json:"resident,omitempty"`
	Nice      *uint64 `json:"nice,omitempty"`
	Ruser     *string `json:"ruser,omitempty"`
	Egroup    *string `json:"egroup,omitempty"`
}

type Ipv6 struct {
	Netmask   *string `json:"netmask,omitempty"`
	Broadcast *string `json:"broadcast,omitempty"`
	Metric    *uint64 `json:"metric,omitempty"`
	Gateway   *string `json:"gateway,omitempty"`
	Dhcp      *string `json:"dhcp,omitempty"`
	Address   *string `json:"address,omitempty"`
}

type Ipv4 struct {
	Address   *string `json:"address,omitempty"`
	Netmask   *string `json:"netmask,omitempty"`
	Broadcast *string `json:"broadcast,omitempty"`
	Metric    *uint64 `json:"metric,omitempty"`
	Gateway   *string `json:"gateway,omitempty"`
	Dhcp      *string `json:"dhcp,omitempty"`
}

type Iface struct {
	Type      *string `json:"type,omitempty"`
	RxBytes   *uint64 `json:"rx_bytes,omitempty"`
	TxPackets *uint64 `json:"tx_packets,omitempty"`
	Ipv6      *Ipv6   `json:"ipv6,omitempty"`
	Name      *string `json:"name,omitempty"`
	State     *string `json:"state,omitempty"`
	Mtu       *uint64 `json:"mtu,omitempty"`
	RxPackets *uint64 `json:"rx_packets,omitempty"`
	Ipv4      *Ipv4   `json:"ipv4,omitempty"`
	Adapter   *string `json:"adapter,omitempty"`
	TxBytes   *uint64 `json:"tx_bytes,omitempty"`
	TxErrors  *uint64 `json:"tx_errors,omitempty"`
	RxErrors  *uint64 `json:"rx_errors,omitempty"`
	TxDropped *uint64 `json:"tx_dropped,omitempty"`
	Mac       *string `json:"mac,omitempty"`
	RxDropped *uint64 `json:"rx_dropped,omitempty"`
}

type Netinfo struct {
	Iface *Iface `json:"iface,omitempty"`
}

type Os struct {
	Version        *string `json:"version,omitempty"`
	Major          *string `json:"major,omitempty"`
	Minor          *string `json:"minor,omitempty"`
	Platform       *string `json:"platform,omitempty"`
	Release        *string `json:"release,omitempty"`
	ReleaseVersion *string `json:"release_version,omitempty"`
	Architecture   *string `json:"architecture,omitempty"`
	Name           *string `json:"name,omitempty"`
	Build          *string `json:"build,omitempty"`
	Sysname        *string `json:"sysname,omitempty"`
	Hostname       *string `json:"hostname,omitempty"`
	Codename       *string `json:"codename,omitempty"`
}

type Port struct {
	RemoteIp   *string `json:"remote_ip,omitempty"`
	TxQueue    *uint64 `json:"tx_queue,omitempty"`
	Inode      *uint64 `json:"inode,omitempty"`
	State      *string `json:"state,omitempty"`
	Pid        *uint64 `json:"pid,omitempty"`
	Process    *string `json:"process,omitempty"`
	LocalPort  *uint64 `json:"local_port,omitempty"`
	LocalIp    *string `json:"local_ip,omitempty"`
	RemotePort *uint64 `json:"remote_port,omitempty"`
	RxQueue    *uint64 `json:"rx_queue,omitempty"`
	Protocol   *string `json:"protocol,omitempty"`
}

type Program struct {
	Version      *string `json:"version,omitempty"`
	Architecture *string `json:"architecture,omitempty"`
	Multiarch    *string `json:"multiarch,omitempty"`
	Source       *string `json:"source,omitempty"`
	Description  *string `json:"description,omitempty"`
	Priority     *string `json:"priority,omitempty"`
	InstallTime  *string `json:"install_time,omitempty"`
	Section      *string `json:"section,omitempty"`
	Size         *uint64 `json:"size,omitempty"`
	Vendor       *string `json:"vendor,omitempty"`
	Location     *string `json:"location,omitempty"`
	Format       *string `json:"format,omitempty"`
	Name         *string `json:"name,omitempty"`
}

type Data struct {
	Audit         *DataAudit     `json:"data_audit,omitempty"`
	Status        *string        `json:"status,omitempty"`
	Uid           *string        `json:"uid,omitempty"`
	Dstip         *string        `json:"dstip,omitempty"`
	Hardware      *Hardware      `json:"hardware,omitempty"`
	Sca           *Sca           `json:"sca,omitempty"`
	Command       *string        `json:"command,omitempty"`
	Timestamp     *int64         `json:"timestamp,omitempty"`
	Title         *string        `json:"title,omitempty"`
	Srcuser       *string        `json:"srcuser,omitempty"`
	ID            *string        `json:"id,omitempty"`
	SystemName    *string        `json:"system_name,omitempty"`
	Oscap         *Oscap         `json:"oscap,omitempty"`
	Dstport       *string        `json:"dstport,omitempty"`
	Dstuser       *string        `json:"dstuser,omitempty"`
	Vulnerability *Vulnerability `json:"vulnerability,omitempty"`
	Aws           *Aws           `json:"aws,omitempty"`
	Action        *string        `json:"action,omitempty"`
	Virustotal    *Virustotal    `json:"virustotal,omitempty"`
	Srcip         *string        `json:"srcip,omitempty"`
	ExtraData     *string        `json:"extra_data,omitempty"`
	Integration   *string        `json:"integration,omitempty"`
	Process       *DataProcess   `json:"data_process,omitempty"`
	Data          *string        `json:"data,omitempty"`
	URL           *string        `json:"url,omitempty"`
	Type          *string        `json:"type,omitempty"`
	Netinfo       *Netinfo       `json:"netinfo,omitempty"`
	Os            *Os            `json:"os,omitempty"`
	Port          *Port          `json:"port,omitempty"`
	Program       *Program       `json:"program,omitempty"`
	Protocol      *string        `json:"protocol,omitempty"`
	Srcport       *string        `json:"srcport,omitempty"`
}

type Alert struct {
	Manager        *Manager     `json:"manager,omitempty"`
	Cluster        *Cluster     `json:"cluster,omitempty"`
	Syscheck       *Syscheck    `json:"syscheck,omitempty"`
	Location       *string      `json:"location,omitempty"`
	Decoder        *Decoder     `json:"decoder,omitempty"`
	Offset         *string      `json:"offset,omitempty"`
	Command        *string      `json:"command,omitempty"`
	Type           *string      `json:"type,omitempty"`
	Timestamp_     *int64       `json:"@timestamp,omitempty"`
	Agent          *Agent       `json:"agent,omitempty"`
	FullLog        *string      `json:"full_log,omitempty"`
	PreviousLog    *string      `json:"previous_log,omitempty"`
	Message        *string      `json:"message,omitempty"`
	Input          *Input       `json:"input,omitempty"`
	Timestamp      *int64       `json:"timestamp,omitempty"`
	Version_       *string      `json:"@version,omitempty"`
	Host           *string      `json:"host,omitempty"`
	Predecoder     *Predecoder  `json:"predecoder,omitempty"`
	ID             *string      `json:"id,omitempty"`
	PreviousOutput *string      `json:"previous_output,omitempty"`
	GeoLocation    *GeoLocation `json:"GeoLocation,omitempty"`
	Rule           *Rule        `json:"rule,omitempty"`
	Data           *Data        `json:"data,omitempty"`
	ProgramName    *string      `json:"program_name,omitempty"`
	Title          *string      `json:"title,omitempty"`
}
