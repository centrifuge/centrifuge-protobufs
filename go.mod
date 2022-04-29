module github.com/centrifuge/centrifuge-protobufs

go 1.18

require (
	github.com/centrifuge/precise-proofs v0.0.0-20190731143435-2ed0cc3986aa
	github.com/golang/protobuf v1.5.0
)

require google.golang.org/protobuf v1.28.0

replace github.com/centrifuge/precise-proofs => github.com/cdamian/precise-proofs v0.0.0-20220429153446-f12245bdfe21
