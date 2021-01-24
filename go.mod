module github.com/fengj3/quicsocks

go 1.15

require (
	github.com/lucas-clemente/quic-go v0.7.1-0.20190825070216-f1d14ecdeafb
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.5.1
)

replace github.com/lucas-clemente/quic-go v0.7.1-0.20190825070216-f1d14ecdeafb => github.com/fbzhong/quic-go v0.7.1-0.20190619145601-64f5a3da04be
