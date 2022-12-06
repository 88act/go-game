/**
 * @Author: 10512203@qq.com
 * @Description:
 * @File: options
 * @Version: 1.0.0
 */

package easysocket

type Options struct {
	certFile string
	keyFile  string
}

type Option func(options *Options)

func newOptions(opt ...Option) Options {
	opts := Options{}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}

func CertFile(v string) Option {
	return func(o *Options) {
		o.certFile = v
	}
}

func KeyFile(v string) Option {
	return func(o *Options) {
		o.keyFile = v
	}
}
