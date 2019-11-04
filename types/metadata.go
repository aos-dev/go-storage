// Code generated by go generate internal/cmd/metdata_gen; DO NOT EDIT.
package types

import (
	"time"
)

// GetAccessKey will get access_key value from metadata.
func (m Metadata) GetAccessKey() (string, bool) {
	v, ok := m[AccessKey]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetAccessKey will set access_key value into metadata.
func (m Metadata) SetAccessKey(v string) {
	m[AccessKey] = v
}

// GetBase will get base value from metadata.
func (m Metadata) GetBase() (string, bool) {
	v, ok := m[Base]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetBase will set base value into metadata.
func (m Metadata) SetBase(v string) {
	m[Base] = v
}

// GetChecksum will get checksum value from metadata.
func (m Metadata) GetChecksum() (string, bool) {
	v, ok := m[Checksum]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetChecksum will set checksum value into metadata.
func (m Metadata) SetChecksum(v string) {
	m[Checksum] = v
}

// GetCount will get count value from metadata.
func (m Metadata) GetCount() (int64, bool) {
	v, ok := m[Count]
	if !ok {
		return 0, false
	}
	return v.(int64), true
}

// SetCount will set count value into metadata.
func (m Metadata) SetCount(v int64) {
	m[Count] = v
}

// GetExpire will get expire value from metadata.
func (m Metadata) GetExpire() (int, bool) {
	v, ok := m[Expire]
	if !ok {
		return 0, false
	}
	return v.(int), true
}

// SetExpire will set expire value into metadata.
func (m Metadata) SetExpire(v int) {
	m[Expire] = v
}

// GetHost will get host value from metadata.
func (m Metadata) GetHost() (string, bool) {
	v, ok := m[Host]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetHost will set host value into metadata.
func (m Metadata) SetHost(v string) {
	m[Host] = v
}

// GetLocation will get location value from metadata.
func (m Metadata) GetLocation() (string, bool) {
	v, ok := m[Location]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetLocation will set location value into metadata.
func (m Metadata) SetLocation(v string) {
	m[Location] = v
}

// GetName will get name value from metadata.
func (m Metadata) GetName() (string, bool) {
	v, ok := m[Name]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetName will set name value into metadata.
func (m Metadata) SetName(v string) {
	m[Name] = v
}

// GetOffset will get offset value from metadata.
func (m Metadata) GetOffset() (int64, bool) {
	v, ok := m[Offset]
	if !ok {
		return 0, false
	}
	return v.(int64), true
}

// SetOffset will set offset value into metadata.
func (m Metadata) SetOffset(v int64) {
	m[Offset] = v
}

// GetPartSize will get part_size value from metadata.
func (m Metadata) GetPartSize() (int64, bool) {
	v, ok := m[PartSize]
	if !ok {
		return 0, false
	}
	return v.(int64), true
}

// SetPartSize will set part_size value into metadata.
func (m Metadata) SetPartSize(v int64) {
	m[PartSize] = v
}

// GetPort will get port value from metadata.
func (m Metadata) GetPort() (int, bool) {
	v, ok := m[Port]
	if !ok {
		return 0, false
	}
	return v.(int), true
}

// SetPort will set port value into metadata.
func (m Metadata) SetPort(v int) {
	m[Port] = v
}

// GetProtocol will get protocol value from metadata.
func (m Metadata) GetProtocol() (string, bool) {
	v, ok := m[Protocol]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetProtocol will set protocol value into metadata.
func (m Metadata) SetProtocol(v string) {
	m[Protocol] = v
}

// GetRecursive will get recursive value from metadata.
func (m Metadata) GetRecursive() (bool, bool) {
	v, ok := m[Recursive]
	if !ok {
		return false, false
	}
	return v.(bool), true
}

// SetRecursive will set recursive value into metadata.
func (m Metadata) SetRecursive(v bool) {
	m[Recursive] = v
}

// GetSecretKey will get secret_key value from metadata.
func (m Metadata) GetSecretKey() (string, bool) {
	v, ok := m[SecretKey]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetSecretKey will set secret_key value into metadata.
func (m Metadata) SetSecretKey(v string) {
	m[SecretKey] = v
}

// GetSize will get size value from metadata.
func (m Metadata) GetSize() (int64, bool) {
	v, ok := m[Size]
	if !ok {
		return 0, false
	}
	return v.(int64), true
}

// SetSize will set size value into metadata.
func (m Metadata) SetSize(v int64) {
	m[Size] = v
}

// GetStorageClass will get storage_class value from metadata.
func (m Metadata) GetStorageClass() (string, bool) {
	v, ok := m[StorageClass]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetStorageClass will set storage_class value into metadata.
func (m Metadata) SetStorageClass(v string) {
	m[StorageClass] = v
}

// GetType will get type value from metadata.
func (m Metadata) GetType() (string, bool) {
	v, ok := m[Type]
	if !ok {
		return "", false
	}
	return v.(string), true
}

// SetType will set type value into metadata.
func (m Metadata) SetType(v string) {
	m[Type] = v
}

// GetUpdatedAt will get updated_at value from metadata.
func (m Metadata) GetUpdatedAt() (time.Time, bool) {
	v, ok := m[UpdatedAt]
	if !ok {
		return time.Time{}, false
	}
	return v.(time.Time), true
}

// SetUpdatedAt will set updated_at value into metadata.
func (m Metadata) SetUpdatedAt(v time.Time) {
	m[UpdatedAt] = v
}
