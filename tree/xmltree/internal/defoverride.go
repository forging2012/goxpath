package defoverride

import (
	"encoding/xml"

	"github.com/ChrisTrenkamp/goxpath/tree"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree/result/xmlattr"
	"github.com/ChrisTrenkamp/goxpath/tree/xmltree/result/xmlele"
)

//DefOverride is the default TreeOverride implementation.  It just builds out
//the data types under github.com/ChrisTrenkamp/goxpath/tree/xmltree/result
type DefOverride struct{}

//Root creates the root document node.  Return your own custom data type here.
func (o DefOverride) Root() tree.Elem {
	root := &xmlele.XMLEle{
		//The GetToken() method should return an empty xml.StartElement struct
		StartElement: xml.StartElement{},
		//The root node does not have any namespaces
		NSStruct: tree.NSStruct{NS: make(map[xml.Name]tree.NS)},
		Attrs:    []*xmlattr.XMLAttr{},
		Children: []tree.Node{},
		Parent:   nil,
	}
	//The parent of the root node must be set to itself, or bad things will happen.
	root.Parent = root

	return root
}

//StartElem appends ele to pos's children.  The child element will be returned.
//ele will already have its fields filled out, they just need to be copied over
//to the new data type.
func (o DefOverride) StartElem(ele *xmlele.XMLEle, pos tree.Elem, dec *xml.Decoder) tree.Elem {
	//When implementing your own StartElem method, pos will be the type of the last call
	//to StartElem, or Root.
	curPos := pos.(*xmlele.XMLEle)

	//Here is where you create your own data type and append to pos's children.
	curPos.Children = append(curPos.Children, ele)
	return ele
}

//AddNode appends n to pos's children.  n's type will be one of *xmlchd.XMLChd, *xmlcomm.XMLComm, *xmlpi.XMLPI.
//Customize the type as needed.
func (o DefOverride) AddNode(n tree.Node, pos tree.Elem, dec *xml.Decoder) {
	curPos := pos.(*xmlele.XMLEle)
	curPos.Children = append(curPos.Children, n)
}

//EndElem marks the end of an element.  The return value is pos's parent.
func (o DefOverride) EndElem(ele xml.EndElement, pos tree.Elem, dec *xml.Decoder) tree.Elem {
	curPos := pos.(*xmlele.XMLEle)
	return curPos.Parent
}