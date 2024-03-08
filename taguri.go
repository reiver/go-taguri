package taguri

// TagURI represents a tag-URI.
//
// Some example tag-URIs include:
//
//	tag:example.com,2024-03-01:note-q744cpgu2s3n7x3
//
//	tag:joeblow@,2024-03-01:0x8b52d56e53ca7a4fa86ae860888c6611
//
//	tag:blogger.com,1999:blog-555
//
//	tag:my-ids.com,2001-09-15:JaneDoe:presentations:UBath2004-05-19
//
// Tag-URIs are a method of creating unique identifiers.
//
// From IETF RFC-4151 —
//
// “A tag is a type of Uniform Resource Identifier (URI) [1] designed to meet the following requirements:”
//
// “1. Identifiers are likely to be unique across space and time, and come from a practically inexhaustible supply.”
//
// “2. Identifiers are relatively convenient for humans to mint (create), read, type, remember etc.”
//
// “3. No central registration is necessary, at least for holders of domain names or email addresses; and there is negligible cost to mint each new identifier.”
//
// “4.  The identifiers are independent of any particular resolution scheme.”
//
type TagURI struct {
	AuthorityName string
	Date          string
	Specific      string
	Fragment      string
}

// String returns the (serialized) tag-URI, based on the data in the 'receiver'.
func (receiver TagURI) String() string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, "tag:"...)
	p = append(p, receiver.AuthorityName...)
	p = append(p, ',')
	p = append(p, receiver.Date...)
	p = append(p, ':')
	p = append(p, receiver.Specific...)

	if "" != receiver.Fragment {
		p = append(p, '#')
		p = append(p, receiver.Fragment...)
	}

	return string(p)
}
