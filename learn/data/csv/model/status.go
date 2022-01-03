package model

type Status int

const (
	Active Status = iota
	Accepted
	Administrative
	Available
	Blended
	Completed
	CustomRole
	Deleted
	DeletedLastCompleted
	Designer
	Inactive
	Observer
	OnCampus
	Online
	Staff
	Student
	StudentOther
	Suspended
	Teacher
	TA
	Error
)

var status = map[string]Status{
	"active":                 Active,
	"accepted":               Accepted,
	"administrative":         Administrative,
	"blended":                Blended,
	"completed":              Completed,
	"costom_role":            CustomRole,
	"deleted":                Deleted,
	"deleted_last_completed": DeletedLastCompleted,
	"designer":               Designer,
	"inactive":               Inactive,
	"observer":               Observer,
	"on_campus":              OnCampus,
	"online":                 Online,
	"staff":                  Staff,
	"student":                Student,
	"student_other":          StudentOther,
	"teacher":                Teacher,
	"ta":                     TA,
	"error":                  Error,
}

func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	case Accepted:
		return "accepted"
	case Administrative:
		return "administrative"
	case Available:
		return "available"
	case Blended:
		return "blended"
	case Completed:
		return "completed"
	case CustomRole:
		return "custom_role"
	case Deleted:
		return "deleted"
	case DeletedLastCompleted:
		return "deleted_last_completed"
	case Designer:
		return "designer"
	case Inactive:
		return "inactive"
	case Observer:
		return "observer"
	case OnCampus:
		return "on_campus"
	case Online:
		return "online"
	case Staff:
		return "staff"
	case Student:
		return "student"
	case StudentOther:
		return "student_other"
	case Suspended:
		return "suspended"
	case Teacher:
		return "teacher"
	case TA:
		return "ta"
	default:
		return "error"
	}
}

func (s Status) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Status) UnmarshalText(data []byte) error {
	v, ok := status[string(data)]
	if !ok {
		return nil
	}

	*s = Status(v)
	return nil
}
