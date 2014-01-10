package tls

import (
    "fmt"
    "reflect"
    "unsafe"
)

func bindata_read(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len
	return b, nil
}

var _assets_tls_trusted_root_crt = "\x2d\x2d\x2d\x2d\x2d\x42\x45\x47\x49\x4e\x20\x43\x45\x52\x54\x49\x46\x49\x43\x41\x54\x45\x2d\x2d\x2d\x2d\x2d\x0a\x4d\x49\x49\x44\x53\x7a\x43\x43\x41\x6a\x4f\x67\x41\x77\x49\x42\x41\x67\x49\x4a\x41\x49\x47\x42\x6d\x6e\x75\x38\x76\x46\x39\x35\x4d\x41\x30\x47\x43\x53\x71\x47\x53\x49\x62\x33\x44\x51\x45\x42\x42\x51\x55\x41\x4d\x44\x77\x78\x43\x7a\x41\x4a\x42\x67\x4e\x56\x0a\x42\x41\x59\x54\x41\x6c\x56\x54\x4d\x52\x4d\x77\x45\x51\x59\x44\x56\x51\x51\x49\x44\x41\x70\x44\x59\x57\x78\x70\x5a\x6d\x39\x79\x62\x6d\x6c\x68\x4d\x52\x67\x77\x46\x67\x59\x44\x56\x51\x51\x4b\x44\x41\x39\x70\x62\x6d\x4e\x76\x62\x6e\x4e\x6f\x63\x6d\x56\x32\x0a\x5a\x57\x46\x69\x62\x47\x55\x77\x48\x68\x63\x4e\x4d\x54\x4d\x78\x4d\x6a\x45\x30\x4d\x6a\x49\x31\x4e\x44\x49\x31\x57\x68\x63\x4e\x4e\x44\x45\x77\x4e\x54\x41\x78\x4d\x6a\x49\x31\x4e\x44\x49\x31\x57\x6a\x41\x38\x4d\x51\x73\x77\x43\x51\x59\x44\x56\x51\x51\x47\x0a\x45\x77\x4a\x56\x55\x7a\x45\x54\x4d\x42\x45\x47\x41\x31\x55\x45\x43\x41\x77\x4b\x51\x32\x46\x73\x61\x57\x5a\x76\x63\x6d\x35\x70\x59\x54\x45\x59\x4d\x42\x59\x47\x41\x31\x55\x45\x43\x67\x77\x50\x61\x57\x35\x6a\x62\x32\x35\x7a\x61\x48\x4a\x6c\x64\x6d\x56\x68\x0a\x59\x6d\x78\x6c\x4d\x49\x49\x42\x49\x6a\x41\x4e\x42\x67\x6b\x71\x68\x6b\x69\x47\x39\x77\x30\x42\x41\x51\x45\x46\x41\x41\x4f\x43\x41\x51\x38\x41\x4d\x49\x49\x42\x43\x67\x4b\x43\x41\x51\x45\x41\x32\x6d\x70\x44\x7a\x65\x4d\x46\x32\x57\x64\x72\x37\x33\x4a\x69\x0a\x5a\x5a\x49\x4d\x6f\x2b\x59\x51\x4e\x62\x4a\x4a\x2b\x33\x31\x73\x42\x32\x2b\x4e\x41\x44\x52\x61\x5a\x59\x6c\x63\x65\x52\x2b\x79\x71\x55\x65\x4f\x6a\x74\x2f\x70\x7a\x35\x4a\x68\x34\x38\x78\x47\x45\x78\x53\x45\x48\x59\x75\x34\x77\x47\x2f\x56\x46\x33\x70\x32\x0a\x2b\x5a\x54\x42\x6a\x53\x30\x6e\x6a\x36\x75\x79\x6e\x44\x36\x30\x56\x31\x75\x32\x62\x70\x75\x50\x63\x61\x66\x66\x4f\x7a\x79\x72\x6a\x49\x6c\x43\x37\x53\x57\x74\x72\x64\x46\x38\x47\x45\x52\x39\x41\x55\x45\x59\x34\x58\x34\x36\x70\x51\x61\x49\x47\x57\x59\x64\x0a\x39\x6e\x32\x2b\x41\x49\x75\x75\x41\x53\x38\x76\x71\x66\x6e\x49\x30\x49\x76\x78\x65\x33\x4f\x7a\x45\x61\x54\x63\x67\x64\x48\x5a\x4e\x55\x6b\x36\x31\x44\x6e\x6c\x4d\x65\x45\x63\x2b\x79\x35\x4f\x65\x66\x37\x33\x63\x4f\x48\x59\x4f\x64\x70\x66\x66\x73\x78\x54\x0a\x68\x73\x79\x73\x42\x76\x30\x45\x67\x79\x63\x31\x30\x68\x74\x69\x6a\x77\x79\x4f\x34\x2f\x2b\x48\x4b\x57\x41\x71\x38\x4e\x39\x50\x30\x31\x50\x6f\x2b\x4c\x79\x63\x48\x67\x42\x65\x61\x34\x2f\x4d\x55\x56\x31\x30\x4f\x53\x45\x77\x37\x32\x35\x52\x39\x4b\x6f\x78\x0a\x57\x61\x41\x34\x39\x63\x45\x63\x70\x36\x64\x4e\x4c\x66\x37\x56\x52\x77\x4d\x5a\x35\x71\x62\x65\x39\x6d\x62\x36\x67\x67\x52\x61\x36\x6e\x6d\x37\x32\x62\x4a\x56\x43\x31\x45\x70\x6b\x45\x66\x51\x6a\x31\x68\x4b\x6c\x39\x4b\x70\x32\x6d\x59\x4e\x65\x42\x6b\x78\x0a\x54\x37\x55\x79\x76\x77\x49\x44\x41\x51\x41\x42\x6f\x31\x41\x77\x54\x6a\x41\x64\x42\x67\x4e\x56\x48\x51\x34\x45\x46\x67\x51\x55\x33\x74\x31\x63\x46\x4e\x71\x45\x51\x33\x4f\x63\x4a\x6f\x5a\x2f\x65\x2f\x72\x49\x4f\x58\x79\x30\x54\x76\x6f\x77\x48\x77\x59\x44\x0a\x56\x52\x30\x6a\x42\x42\x67\x77\x46\x6f\x41\x55\x33\x74\x31\x63\x46\x4e\x71\x45\x51\x33\x4f\x63\x4a\x6f\x5a\x2f\x65\x2f\x72\x49\x4f\x58\x79\x30\x54\x76\x6f\x77\x44\x41\x59\x44\x56\x52\x30\x54\x42\x41\x55\x77\x41\x77\x45\x42\x2f\x7a\x41\x4e\x42\x67\x6b\x71\x0a\x68\x6b\x69\x47\x39\x77\x30\x42\x41\x51\x55\x46\x41\x41\x4f\x43\x41\x51\x45\x41\x68\x7a\x6f\x71\x79\x57\x4c\x37\x4f\x2f\x71\x79\x73\x54\x63\x44\x37\x51\x43\x65\x43\x6d\x33\x4d\x46\x65\x41\x65\x41\x72\x68\x75\x35\x6a\x2f\x6b\x52\x6c\x58\x36\x6d\x74\x45\x49\x0a\x73\x45\x2b\x75\x52\x6b\x49\x71\x48\x4d\x49\x69\x71\x69\x48\x49\x76\x65\x52\x30\x76\x39\x41\x73\x55\x56\x70\x46\x34\x77\x6f\x51\x49\x70\x42\x41\x35\x53\x6c\x79\x54\x56\x38\x38\x41\x41\x61\x6c\x6b\x6e\x75\x61\x4d\x73\x48\x50\x5a\x6e\x77\x30\x6e\x6f\x4f\x68\x0a\x53\x54\x69\x58\x4c\x35\x2b\x4e\x57\x56\x36\x6c\x54\x45\x30\x74\x6b\x75\x51\x6e\x4a\x39\x35\x69\x79\x4e\x30\x31\x65\x47\x2f\x70\x53\x2f\x65\x78\x4e\x2b\x6f\x54\x77\x49\x76\x69\x5a\x78\x47\x75\x70\x2b\x5a\x4e\x30\x59\x4c\x57\x4e\x46\x6c\x71\x77\x2f\x46\x4f\x0a\x33\x59\x4a\x65\x72\x46\x46\x6c\x35\x53\x68\x47\x70\x49\x62\x45\x55\x6e\x59\x4b\x6a\x32\x69\x47\x50\x70\x65\x72\x31\x39\x32\x73\x31\x4d\x7a\x36\x67\x69\x73\x63\x6b\x50\x67\x4f\x69\x71\x49\x4a\x6f\x50\x5a\x30\x65\x4c\x4b\x6e\x5a\x2f\x41\x43\x4e\x63\x7a\x54\x0a\x69\x48\x31\x43\x6b\x46\x70\x57\x65\x62\x5a\x2f\x52\x71\x31\x73\x42\x46\x66\x58\x6e\x6e\x6f\x4a\x44\x35\x58\x45\x62\x2b\x38\x50\x5a\x77\x32\x65\x49\x48\x4e\x67\x49\x62\x48\x55\x66\x61\x34\x42\x42\x74\x49\x73\x78\x72\x7a\x53\x58\x64\x77\x63\x4b\x6b\x57\x64\x0a\x58\x79\x71\x72\x51\x6f\x44\x66\x77\x4d\x6e\x34\x4b\x31\x71\x65\x30\x73\x66\x50\x4b\x66\x79\x47\x73\x35\x54\x74\x61\x72\x6d\x62\x77\x4d\x50\x70\x54\x59\x73\x64\x42\x51\x3d\x3d\x0a\x2d\x2d\x2d\x2d\x2d\x45\x4e\x44\x20\x43\x45\x52\x54\x49\x46\x49\x43\x41\x54\x45\x2d\x2d\x2d\x2d\x2d\x0a"

func assets_tls_trusted_root_crt() ([]byte, error) {
	return bindata_read(
		_assets_tls_trusted_root_crt,
		"assets/tls/trusted.root.crt",
	)
}

var _assets_tls_snakeoil_root_crt = "\x2d\x2d\x2d\x2d\x2d\x42\x45\x47\x49\x4e\x20\x43\x45\x52\x54\x49\x46\x49\x43\x41\x54\x45\x2d\x2d\x2d\x2d\x2d\x0a\x4d\x49\x49\x46\x46\x44\x43\x43\x41\x76\x77\x43\x43\x51\x43\x6b\x62\x4e\x30\x52\x47\x2f\x6f\x31\x35\x44\x41\x4e\x42\x67\x6b\x71\x68\x6b\x69\x47\x39\x77\x30\x42\x41\x51\x55\x46\x41\x44\x42\x4d\x4d\x51\x73\x77\x43\x51\x59\x44\x56\x51\x51\x47\x45\x77\x4a\x56\x0a\x55\x7a\x45\x54\x4d\x42\x45\x47\x41\x31\x55\x45\x43\x42\x4d\x4b\x51\x32\x46\x73\x61\x57\x5a\x76\x63\x6d\x35\x70\x59\x54\x45\x53\x4d\x42\x41\x47\x41\x31\x55\x45\x43\x68\x4d\x4a\x62\x6d\x64\x79\x62\x32\x73\x75\x59\x32\x39\x74\x4d\x52\x51\x77\x45\x67\x59\x44\x0a\x56\x51\x51\x44\x46\x41\x73\x71\x4c\x6d\x35\x6e\x63\x6d\x39\x72\x4c\x6d\x4e\x76\x62\x54\x41\x65\x46\x77\x30\x78\x4d\x7a\x41\x32\x4d\x44\x4d\x77\x4d\x7a\x55\x78\x4e\x54\x5a\x61\x46\x77\x30\x79\x4d\x7a\x41\x32\x4d\x44\x45\x77\x4d\x7a\x55\x78\x4e\x54\x5a\x61\x0a\x4d\x45\x77\x78\x43\x7a\x41\x4a\x42\x67\x4e\x56\x42\x41\x59\x54\x41\x6c\x56\x54\x4d\x52\x4d\x77\x45\x51\x59\x44\x56\x51\x51\x49\x45\x77\x70\x44\x59\x57\x78\x70\x5a\x6d\x39\x79\x62\x6d\x6c\x68\x4d\x52\x49\x77\x45\x41\x59\x44\x56\x51\x51\x4b\x45\x77\x6c\x75\x0a\x5a\x33\x4a\x76\x61\x79\x35\x6a\x62\x32\x30\x78\x46\x44\x41\x53\x42\x67\x4e\x56\x42\x41\x4d\x55\x43\x79\x6f\x75\x62\x6d\x64\x79\x62\x32\x73\x75\x59\x32\x39\x74\x4d\x49\x49\x43\x49\x6a\x41\x4e\x42\x67\x6b\x71\x68\x6b\x69\x47\x39\x77\x30\x42\x41\x51\x45\x46\x0a\x41\x41\x4f\x43\x41\x67\x38\x41\x4d\x49\x49\x43\x43\x67\x4b\x43\x41\x67\x45\x41\x36\x51\x72\x79\x58\x65\x4b\x6c\x38\x41\x57\x57\x61\x39\x75\x47\x32\x55\x62\x53\x4f\x70\x6f\x6f\x48\x37\x34\x7a\x4c\x6b\x58\x73\x33\x46\x5a\x66\x6b\x39\x67\x4b\x76\x71\x6b\x69\x0a\x7a\x58\x58\x51\x43\x52\x6d\x74\x55\x36\x44\x79\x6e\x30\x2b\x4f\x75\x53\x33\x73\x45\x2f\x72\x52\x6d\x5a\x41\x73\x53\x6a\x6b\x51\x47\x2f\x59\x44\x74\x64\x45\x2f\x53\x67\x4c\x34\x64\x56\x36\x53\x36\x32\x71\x69\x51\x6e\x67\x50\x6f\x6b\x6a\x52\x30\x55\x53\x68\x0a\x50\x43\x34\x48\x77\x62\x38\x54\x6a\x4d\x39\x57\x35\x43\x64\x2b\x6f\x77\x56\x7a\x4d\x51\x30\x76\x6c\x30\x41\x59\x68\x51\x6b\x38\x59\x63\x2f\x30\x76\x58\x2b\x7a\x44\x4f\x77\x6d\x52\x57\x47\x6a\x4e\x4b\x50\x71\x34\x32\x32\x75\x73\x46\x39\x43\x4a\x46\x63\x2f\x0a\x38\x51\x59\x2b\x4f\x44\x4a\x44\x48\x75\x6e\x38\x56\x56\x41\x6b\x71\x33\x58\x66\x63\x50\x58\x67\x79\x74\x48\x49\x71\x78\x76\x53\x4a\x6e\x59\x67\x44\x6f\x75\x46\x43\x41\x2b\x47\x54\x73\x4b\x70\x2f\x36\x35\x53\x35\x63\x69\x67\x53\x6c\x49\x72\x51\x5a\x62\x48\x0a\x37\x37\x35\x63\x54\x57\x68\x43\x6a\x76\x59\x6e\x71\x36\x67\x7a\x79\x72\x6b\x33\x52\x69\x47\x64\x62\x31\x49\x47\x75\x49\x4a\x66\x74\x4d\x4a\x78\x75\x4a\x79\x4a\x56\x62\x66\x54\x46\x74\x71\x67\x4d\x47\x54\x6d\x6a\x48\x5a\x78\x69\x4c\x76\x4d\x37\x64\x7a\x37\x0a\x6a\x2f\x62\x6d\x72\x7a\x34\x50\x76\x6e\x68\x62\x51\x53\x5a\x5a\x4c\x68\x73\x76\x50\x31\x6f\x38\x6d\x78\x6e\x6f\x4e\x4d\x70\x6f\x2f\x54\x6f\x35\x74\x48\x70\x2f\x54\x73\x36\x62\x35\x46\x51\x4e\x4c\x37\x46\x48\x70\x6d\x4f\x56\x4c\x41\x6f\x51\x33\x46\x64\x58\x0a\x56\x72\x79\x54\x6a\x6f\x53\x6a\x69\x45\x33\x4a\x4c\x44\x47\x5a\x49\x4e\x51\x39\x4d\x46\x45\x50\x67\x50\x7a\x52\x38\x6d\x72\x7a\x71\x46\x6f\x2f\x36\x65\x37\x75\x42\x34\x41\x59\x6c\x4b\x6f\x51\x6f\x30\x31\x4b\x7a\x78\x2b\x59\x6d\x56\x77\x52\x4b\x45\x74\x72\x0a\x56\x43\x54\x52\x5a\x52\x63\x6c\x36\x36\x2b\x67\x4d\x6b\x63\x58\x30\x72\x79\x6f\x56\x6e\x67\x67\x6a\x49\x57\x57\x75\x34\x64\x38\x75\x41\x68\x33\x6a\x66\x2b\x2b\x4b\x64\x30\x44\x6a\x62\x2f\x6c\x37\x44\x43\x50\x70\x45\x67\x53\x4a\x77\x5a\x59\x54\x6a\x43\x4c\x0a\x5a\x36\x53\x78\x69\x42\x77\x51\x34\x6f\x39\x64\x45\x51\x61\x64\x4d\x47\x67\x6b\x33\x74\x6c\x44\x46\x43\x42\x73\x72\x48\x6f\x71\x37\x4e\x79\x7a\x76\x58\x48\x50\x30\x42\x46\x32\x48\x4b\x62\x38\x4b\x52\x45\x42\x45\x4b\x43\x49\x75\x51\x6a\x39\x52\x43\x68\x57\x0a\x67\x30\x37\x7a\x6d\x4f\x70\x6a\x6e\x67\x57\x73\x30\x43\x58\x61\x59\x6c\x79\x2b\x54\x44\x50\x2b\x35\x44\x5a\x43\x4d\x47\x44\x39\x6b\x6d\x58\x6b\x51\x59\x39\x71\x2f\x7a\x71\x71\x76\x4d\x74\x2b\x54\x2b\x2f\x54\x42\x4b\x39\x6c\x77\x55\x73\x6f\x69\x32\x55\x63\x0a\x76\x39\x33\x77\x53\x2b\x54\x4e\x75\x30\x36\x61\x52\x6f\x70\x71\x50\x6f\x39\x59\x5a\x72\x33\x38\x6b\x61\x33\x78\x4b\x50\x69\x4f\x39\x36\x34\x70\x6b\x32\x42\x6f\x46\x4e\x35\x37\x67\x37\x36\x37\x47\x38\x6b\x39\x54\x62\x68\x6b\x42\x78\x69\x74\x76\x46\x55\x43\x0a\x41\x77\x45\x41\x41\x54\x41\x4e\x42\x67\x6b\x71\x68\x6b\x69\x47\x39\x77\x30\x42\x41\x51\x55\x46\x41\x41\x4f\x43\x41\x67\x45\x41\x52\x66\x2f\x68\x56\x59\x6e\x74\x7a\x55\x77\x46\x55\x67\x51\x72\x57\x44\x30\x6c\x2f\x55\x61\x43\x42\x67\x72\x6c\x76\x78\x56\x43\x0a\x79\x55\x61\x38\x49\x73\x6a\x33\x76\x65\x7a\x41\x68\x46\x53\x79\x5a\x6e\x74\x45\x4c\x2b\x45\x4c\x46\x76\x38\x76\x76\x51\x62\x74\x42\x47\x48\x48\x2f\x74\x43\x6e\x37\x36\x57\x75\x71\x6a\x77\x4f\x6a\x56\x4c\x32\x33\x79\x78\x6b\x61\x4a\x73\x72\x6e\x52\x39\x2b\x0a\x54\x52\x4e\x46\x6e\x56\x65\x42\x38\x31\x35\x37\x2b\x49\x46\x36\x48\x4b\x7a\x4c\x43\x4c\x2f\x48\x49\x41\x69\x51\x30\x6b\x77\x2f\x32\x4f\x53\x4c\x44\x30\x6c\x5a\x6e\x41\x6b\x67\x32\x34\x41\x30\x2f\x39\x53\x48\x63\x70\x49\x35\x47\x41\x30\x57\x6c\x54\x68\x45\x0a\x34\x47\x71\x67\x63\x55\x69\x4e\x39\x6d\x2b\x6d\x4c\x38\x6a\x57\x47\x33\x67\x6a\x2b\x53\x58\x43\x37\x49\x63\x56\x53\x33\x76\x41\x76\x53\x31\x4a\x37\x4b\x7a\x30\x4e\x7a\x54\x68\x31\x64\x59\x6b\x51\x4e\x57\x4a\x6c\x61\x75\x4f\x32\x51\x6e\x39\x35\x54\x39\x39\x0a\x70\x6c\x6b\x50\x50\x68\x38\x37\x79\x5a\x4f\x39\x61\x39\x62\x78\x70\x58\x39\x50\x55\x4a\x6b\x54\x4a\x7a\x4f\x77\x55\x6b\x5a\x49\x53\x52\x5a\x45\x62\x54\x41\x30\x43\x66\x73\x70\x55\x70\x71\x2f\x70\x68\x7a\x75\x54\x56\x69\x79\x37\x6f\x32\x66\x72\x2b\x54\x6f\x0a\x78\x56\x56\x61\x32\x61\x4b\x54\x39\x31\x32\x76\x6c\x51\x61\x64\x57\x77\x35\x6f\x71\x45\x4b\x37\x78\x79\x78\x54\x71\x73\x59\x77\x56\x37\x43\x55\x6c\x6a\x74\x43\x6e\x68\x70\x53\x37\x77\x4f\x5a\x68\x77\x7a\x49\x33\x71\x55\x6b\x39\x48\x4b\x48\x30\x52\x74\x39\x0a\x2f\x48\x51\x73\x41\x4e\x75\x53\x69\x6b\x5a\x6f\x73\x4e\x76\x64\x4d\x33\x2f\x68\x76\x33\x63\x35\x44\x52\x55\x4f\x77\x4b\x69\x4b\x64\x62\x67\x5a\x43\x79\x71\x51\x66\x38\x58\x53\x65\x4a\x52\x4d\x31\x69\x51\x76\x63\x48\x6f\x38\x55\x38\x6b\x45\x4a\x45\x45\x74\x0a\x64\x6d\x66\x74\x6e\x2b\x30\x67\x48\x33\x52\x50\x73\x56\x30\x32\x38\x2b\x37\x58\x44\x36\x77\x72\x61\x71\x66\x63\x37\x64\x7a\x4e\x7a\x4c\x6e\x71\x32\x72\x44\x4c\x53\x41\x66\x47\x33\x54\x37\x70\x70\x30\x6e\x31\x39\x4a\x79\x55\x69\x65\x57\x73\x54\x52\x34\x35\x0a\x70\x47\x63\x77\x4e\x70\x58\x44\x6b\x2b\x77\x65\x71\x51\x4b\x6b\x6e\x77\x62\x4e\x6f\x68\x61\x33\x6f\x35\x31\x58\x71\x2f\x31\x6e\x68\x52\x74\x71\x72\x55\x50\x49\x45\x67\x4f\x6e\x42\x57\x42\x43\x56\x38\x76\x6b\x68\x58\x4d\x68\x39\x56\x62\x65\x2b\x6f\x41\x6d\x0a\x4c\x6d\x77\x65\x4e\x31\x4d\x72\x36\x4d\x6a\x72\x48\x57\x64\x64\x56\x6e\x2b\x4a\x47\x63\x42\x35\x70\x2b\x41\x4d\x57\x72\x38\x7a\x45\x2b\x62\x68\x70\x50\x43\x41\x6e\x75\x70\x46\x52\x39\x38\x7a\x33\x66\x62\x4f\x6c\x65\x4d\x43\x57\x71\x36\x51\x2b\x68\x4d\x50\x0a\x4d\x55\x54\x68\x71\x4a\x52\x4f\x48\x61\x6d\x48\x52\x49\x4a\x33\x49\x7a\x38\x77\x68\x49\x72\x6a\x37\x45\x4b\x6b\x44\x42\x4b\x4c\x45\x66\x45\x31\x45\x78\x53\x33\x42\x39\x56\x51\x2b\x59\x64\x48\x79\x36\x73\x6a\x65\x78\x64\x6d\x43\x49\x51\x67\x45\x6c\x71\x34\x0a\x34\x53\x4d\x75\x59\x2f\x4a\x6b\x5a\x6c\x6f\x3d\x0a\x2d\x2d\x2d\x2d\x2d\x45\x4e\x44\x20\x43\x45\x52\x54\x49\x46\x49\x43\x41\x54\x45\x2d\x2d\x2d\x2d\x2d\x0a"

func assets_tls_snakeoil_root_crt() ([]byte, error) {
	return bindata_read(
		_assets_tls_snakeoil_root_crt,
		"assets/tls/snakeoil.root.crt",
	)
}

var _assets_tls_snakeoil_key = "\x2d\x2d\x2d\x2d\x2d\x42\x45\x47\x49\x4e\x20\x52\x53\x41\x20\x50\x52\x49\x56\x41\x54\x45\x20\x4b\x45\x59\x2d\x2d\x2d\x2d\x2d\x0a\x4d\x49\x49\x4a\x4b\x41\x49\x42\x41\x41\x4b\x43\x41\x67\x45\x41\x32\x69\x35\x63\x4c\x72\x41\x6c\x70\x77\x38\x67\x56\x67\x53\x49\x6f\x67\x59\x61\x47\x57\x4a\x4b\x48\x47\x7a\x4c\x58\x7a\x55\x73\x6e\x54\x72\x68\x6c\x4d\x50\x63\x42\x45\x78\x4f\x4c\x30\x34\x30\x0a\x36\x78\x47\x51\x64\x64\x4f\x67\x7a\x51\x4a\x58\x56\x6b\x30\x4f\x65\x49\x2b\x6e\x63\x6b\x67\x75\x49\x59\x76\x6e\x59\x55\x2b\x6e\x53\x46\x30\x6c\x36\x42\x44\x36\x6e\x50\x50\x2b\x41\x4f\x4e\x66\x53\x48\x30\x77\x4a\x5a\x37\x45\x7a\x69\x79\x6f\x75\x6c\x6a\x52\x0a\x4e\x38\x53\x71\x54\x32\x55\x64\x59\x6f\x38\x6d\x44\x34\x34\x4d\x72\x7a\x42\x72\x58\x2f\x43\x35\x41\x34\x72\x37\x36\x62\x39\x64\x6a\x66\x42\x78\x74\x6d\x41\x75\x50\x62\x63\x51\x64\x77\x62\x4a\x75\x58\x2f\x45\x59\x6d\x52\x70\x35\x71\x6f\x59\x32\x2b\x73\x6d\x0a\x6a\x2f\x68\x47\x42\x49\x76\x2b\x4e\x4f\x6d\x6e\x52\x30\x4e\x59\x70\x53\x59\x6f\x47\x6c\x4d\x66\x67\x2b\x34\x50\x70\x75\x53\x58\x79\x6f\x70\x41\x41\x2f\x56\x36\x36\x58\x70\x6c\x4c\x34\x61\x76\x2b\x57\x55\x56\x67\x6c\x61\x72\x43\x42\x33\x64\x49\x50\x64\x61\x0a\x48\x2f\x72\x56\x6a\x69\x66\x71\x77\x4f\x36\x78\x36\x62\x59\x31\x43\x57\x32\x65\x55\x58\x42\x37\x6c\x56\x4c\x55\x74\x2b\x69\x6c\x57\x30\x37\x64\x37\x59\x33\x6c\x2b\x39\x34\x52\x52\x31\x75\x44\x5a\x67\x4f\x7a\x74\x36\x38\x31\x68\x55\x72\x37\x57\x4c\x53\x53\x0a\x4f\x6a\x52\x30\x50\x63\x4f\x31\x75\x61\x46\x37\x58\x63\x35\x44\x39\x74\x71\x35\x2f\x46\x77\x64\x7a\x58\x47\x4f\x30\x61\x77\x6e\x6e\x76\x43\x7a\x79\x64\x34\x38\x76\x68\x36\x6a\x36\x38\x57\x34\x50\x6c\x43\x4e\x7a\x30\x52\x4a\x2f\x4f\x36\x64\x4a\x36\x62\x53\x0a\x45\x4c\x6b\x41\x6d\x41\x57\x75\x78\x62\x4d\x68\x69\x79\x76\x74\x53\x79\x4b\x34\x6f\x62\x69\x42\x42\x30\x50\x58\x73\x76\x34\x4f\x69\x42\x55\x4b\x70\x46\x78\x31\x68\x65\x36\x43\x73\x38\x51\x6b\x34\x32\x4c\x78\x67\x62\x64\x35\x46\x45\x2f\x39\x44\x32\x53\x56\x0a\x66\x67\x39\x62\x4c\x73\x35\x6d\x50\x73\x70\x66\x2f\x45\x47\x75\x39\x4c\x6b\x5a\x34\x4d\x32\x71\x73\x4e\x57\x77\x38\x76\x52\x62\x62\x67\x55\x6e\x2f\x39\x4b\x36\x7a\x73\x46\x49\x58\x52\x39\x7a\x6b\x4e\x4a\x69\x71\x30\x50\x68\x6d\x42\x2f\x55\x5a\x47\x4c\x54\x0a\x37\x46\x4a\x72\x65\x62\x30\x43\x41\x6e\x42\x59\x34\x65\x62\x71\x49\x4c\x53\x59\x66\x42\x31\x32\x6a\x4d\x38\x74\x33\x34\x49\x2f\x38\x7a\x73\x39\x6a\x4d\x39\x45\x4f\x4a\x5a\x6a\x39\x36\x47\x67\x49\x57\x4b\x70\x45\x43\x43\x56\x67\x37\x53\x61\x45\x72\x49\x6d\x0a\x2b\x68\x4d\x36\x4f\x72\x70\x4d\x46\x68\x69\x6e\x2f\x77\x74\x78\x75\x70\x2b\x70\x53\x51\x6a\x6f\x4e\x4b\x64\x6e\x5a\x38\x31\x31\x63\x6f\x50\x4b\x4b\x73\x78\x49\x35\x50\x67\x72\x57\x36\x44\x55\x77\x64\x71\x72\x79\x68\x75\x35\x77\x75\x72\x48\x37\x6b\x65\x4a\x0a\x37\x6c\x33\x6f\x49\x70\x47\x43\x63\x33\x58\x6d\x77\x6e\x41\x59\x6f\x6a\x4c\x62\x74\x2b\x44\x6f\x73\x44\x4f\x4c\x55\x5a\x61\x34\x4c\x4e\x6e\x57\x51\x65\x71\x4f\x35\x43\x48\x65\x68\x56\x2b\x34\x45\x71\x63\x56\x2b\x78\x30\x74\x53\x63\x55\x43\x41\x77\x45\x41\x0a\x41\x51\x4b\x43\x41\x67\x42\x62\x36\x52\x75\x38\x4c\x30\x67\x74\x55\x42\x6e\x33\x49\x6f\x48\x4d\x66\x33\x57\x50\x4b\x2f\x43\x38\x65\x4c\x68\x54\x71\x7a\x72\x59\x49\x57\x33\x57\x46\x59\x77\x68\x34\x32\x4d\x73\x57\x6d\x33\x41\x65\x4f\x32\x36\x4e\x53\x53\x51\x0a\x4f\x47\x52\x43\x58\x73\x4f\x78\x31\x68\x4a\x62\x2b\x6a\x77\x30\x74\x5a\x4d\x4c\x55\x31\x72\x4e\x43\x54\x42\x6d\x79\x6f\x42\x49\x6a\x69\x42\x36\x6a\x30\x34\x63\x59\x32\x42\x63\x2b\x4c\x30\x2f\x66\x57\x43\x32\x33\x36\x4f\x44\x4d\x72\x33\x73\x4a\x46\x52\x30\x0a\x71\x49\x6b\x49\x46\x48\x63\x54\x64\x66\x70\x46\x75\x45\x71\x34\x53\x31\x78\x44\x34\x2f\x47\x74\x55\x5a\x55\x56\x6c\x76\x37\x6a\x30\x4c\x4b\x47\x38\x62\x30\x59\x2f\x39\x48\x6a\x41\x52\x6e\x37\x71\x62\x77\x2f\x4b\x4a\x68\x65\x48\x65\x43\x68\x47\x62\x68\x45\x0a\x34\x67\x6b\x74\x35\x42\x6a\x37\x75\x55\x38\x37\x68\x37\x6a\x48\x41\x77\x70\x6b\x36\x2f\x64\x6c\x77\x30\x65\x6b\x59\x30\x30\x62\x2f\x67\x75\x53\x4d\x64\x4c\x2f\x35\x4b\x31\x69\x38\x73\x2b\x70\x34\x36\x71\x37\x73\x48\x65\x75\x38\x53\x50\x31\x64\x71\x74\x57\x0a\x43\x7a\x65\x33\x6c\x4b\x4a\x54\x44\x6e\x4b\x62\x4c\x42\x39\x6a\x6b\x44\x6b\x38\x49\x43\x31\x49\x67\x62\x6a\x4c\x30\x66\x4d\x49\x58\x30\x77\x34\x47\x7a\x30\x48\x52\x4a\x66\x34\x30\x54\x35\x69\x6f\x47\x75\x78\x75\x70\x2b\x2f\x46\x55\x6e\x43\x6d\x79\x64\x36\x0a\x77\x36\x51\x4d\x71\x45\x2f\x4a\x4e\x65\x73\x54\x66\x46\x71\x78\x71\x52\x7a\x5a\x42\x77\x54\x4a\x31\x2b\x78\x6b\x58\x6b\x53\x46\x6d\x46\x68\x5a\x39\x4d\x63\x66\x78\x6d\x45\x74\x6f\x7a\x61\x75\x30\x4c\x48\x36\x54\x74\x6b\x47\x6d\x48\x76\x66\x6d\x6e\x36\x2f\x0a\x49\x38\x49\x6d\x6c\x49\x52\x79\x48\x47\x4f\x55\x56\x48\x73\x45\x41\x52\x38\x63\x37\x61\x44\x39\x64\x61\x67\x30\x4c\x62\x58\x6d\x6a\x79\x6f\x6c\x66\x69\x6e\x38\x44\x6f\x77\x31\x33\x30\x2f\x47\x33\x68\x54\x58\x59\x6e\x44\x45\x54\x78\x36\x36\x6b\x31\x52\x54\x0a\x42\x68\x36\x6b\x44\x45\x59\x76\x75\x78\x76\x46\x71\x6b\x58\x59\x4b\x52\x41\x75\x35\x75\x31\x31\x76\x51\x41\x39\x75\x64\x59\x71\x4e\x34\x69\x2b\x68\x74\x51\x32\x4f\x4f\x59\x4f\x45\x4b\x4e\x58\x38\x4b\x78\x79\x35\x36\x32\x72\x58\x67\x64\x38\x35\x4d\x79\x54\x0a\x65\x54\x4c\x34\x44\x6d\x68\x2f\x36\x4a\x37\x56\x74\x57\x47\x34\x6e\x4e\x32\x47\x38\x70\x6e\x6a\x35\x4c\x67\x39\x39\x51\x38\x66\x7a\x45\x46\x51\x6d\x46\x4c\x46\x63\x4b\x30\x6b\x4a\x77\x48\x75\x30\x34\x70\x4a\x61\x71\x72\x48\x4c\x2b\x75\x48\x48\x64\x35\x54\x0a\x4b\x4f\x66\x50\x41\x72\x30\x70\x55\x62\x65\x75\x32\x73\x56\x64\x39\x63\x68\x67\x44\x68\x41\x44\x55\x45\x31\x2f\x38\x59\x50\x68\x53\x43\x73\x48\x75\x64\x4d\x34\x4e\x6f\x34\x74\x44\x43\x5a\x56\x45\x71\x4b\x4a\x31\x70\x47\x4c\x43\x70\x55\x77\x33\x48\x31\x76\x0a\x68\x30\x61\x41\x2f\x42\x35\x38\x32\x70\x32\x43\x47\x46\x70\x47\x6f\x49\x6e\x4b\x66\x46\x4b\x76\x55\x62\x59\x74\x32\x6e\x65\x44\x5a\x5a\x51\x79\x67\x6d\x59\x43\x55\x48\x79\x77\x71\x4e\x78\x51\x41\x51\x4b\x43\x41\x51\x45\x41\x2f\x48\x30\x4d\x43\x4a\x59\x56\x0a\x56\x52\x52\x4c\x75\x66\x5a\x6b\x7a\x33\x66\x70\x43\x2b\x77\x54\x2b\x63\x6f\x71\x62\x67\x75\x77\x72\x54\x32\x62\x70\x56\x42\x75\x4e\x64\x76\x79\x44\x4b\x59\x49\x6b\x62\x63\x53\x44\x74\x49\x56\x4e\x73\x4f\x2f\x31\x70\x31\x2f\x43\x47\x6f\x72\x6d\x44\x6c\x6c\x0a\x6b\x31\x6a\x4e\x46\x68\x52\x2b\x6c\x71\x42\x42\x38\x51\x51\x66\x62\x31\x5a\x65\x74\x50\x6f\x68\x78\x7a\x50\x4e\x6e\x63\x50\x69\x4a\x79\x70\x62\x39\x55\x69\x76\x75\x37\x47\x47\x31\x4d\x4e\x55\x48\x4a\x55\x65\x63\x46\x74\x38\x75\x49\x54\x63\x59\x39\x33\x73\x0a\x55\x67\x52\x42\x56\x73\x43\x70\x4c\x67\x55\x50\x37\x53\x46\x36\x6d\x31\x4d\x70\x75\x32\x42\x61\x58\x51\x2f\x73\x45\x4d\x38\x64\x59\x43\x53\x6d\x74\x34\x46\x61\x78\x37\x4c\x57\x30\x30\x63\x49\x70\x4c\x6c\x65\x5a\x6d\x43\x58\x6f\x63\x2f\x6a\x73\x54\x55\x4d\x0a\x69\x78\x59\x54\x4d\x35\x72\x35\x55\x57\x58\x42\x68\x6c\x31\x39\x36\x4c\x66\x75\x2b\x57\x6a\x54\x6d\x6f\x78\x47\x67\x46\x65\x54\x77\x46\x38\x78\x31\x77\x75\x6f\x37\x6d\x51\x76\x50\x66\x72\x69\x70\x76\x55\x58\x2f\x62\x6a\x47\x52\x34\x6f\x35\x4a\x51\x4f\x77\x0a\x31\x49\x30\x58\x4b\x31\x33\x79\x72\x68\x46\x6b\x45\x41\x49\x56\x36\x31\x6e\x48\x42\x47\x69\x79\x68\x6d\x4a\x53\x68\x50\x61\x53\x49\x32\x75\x44\x75\x64\x78\x30\x5a\x6b\x66\x75\x55\x49\x4d\x31\x55\x46\x76\x71\x65\x71\x57\x73\x6a\x41\x54\x61\x4c\x6d\x4f\x57\x0a\x58\x37\x6c\x45\x30\x71\x4e\x67\x37\x35\x6c\x73\x52\x51\x4b\x43\x41\x51\x45\x41\x33\x54\x63\x71\x67\x76\x69\x73\x73\x39\x61\x34\x4f\x77\x72\x30\x42\x61\x67\x6c\x65\x4f\x62\x55\x6e\x54\x4d\x46\x7a\x6b\x38\x62\x67\x70\x65\x6d\x4e\x64\x7a\x69\x34\x37\x65\x47\x0a\x37\x43\x35\x37\x32\x57\x78\x38\x47\x73\x36\x58\x67\x4e\x62\x4f\x32\x38\x6c\x70\x45\x30\x50\x44\x37\x4b\x55\x2f\x52\x76\x74\x33\x36\x77\x50\x39\x6d\x7a\x47\x69\x50\x30\x4d\x33\x5a\x72\x34\x68\x38\x31\x41\x66\x4c\x76\x66\x74\x6f\x4d\x4d\x73\x36\x79\x74\x6b\x0a\x77\x48\x4a\x4f\x64\x74\x50\x55\x43\x5a\x71\x76\x6e\x62\x6c\x76\x4b\x4f\x49\x34\x30\x39\x6e\x47\x6e\x6f\x55\x44\x2b\x33\x6e\x66\x48\x2f\x71\x71\x48\x6b\x34\x5a\x4c\x33\x66\x6e\x34\x76\x48\x62\x78\x77\x4d\x6f\x64\x5a\x66\x4f\x58\x52\x72\x48\x32\x68\x4a\x72\x0a\x45\x56\x71\x41\x6e\x6e\x65\x78\x69\x39\x53\x30\x4a\x62\x4a\x4a\x78\x64\x63\x53\x6c\x42\x4d\x49\x35\x76\x51\x4d\x4b\x47\x4c\x43\x6a\x35\x32\x43\x32\x6d\x66\x63\x46\x2b\x58\x52\x4d\x75\x46\x38\x38\x62\x77\x66\x4a\x37\x58\x63\x62\x33\x79\x68\x79\x76\x38\x69\x0a\x70\x58\x48\x48\x5a\x5a\x39\x51\x4b\x45\x37\x65\x47\x37\x36\x55\x4f\x2b\x45\x47\x4b\x7a\x38\x67\x6e\x33\x77\x57\x4b\x75\x68\x35\x48\x58\x70\x33\x48\x6f\x79\x4b\x69\x64\x67\x7a\x58\x74\x50\x49\x62\x75\x38\x66\x64\x34\x55\x7a\x58\x6c\x2f\x42\x54\x4d\x46\x63\x0a\x6a\x57\x41\x47\x30\x59\x63\x73\x6d\x35\x64\x54\x6f\x62\x51\x30\x59\x6f\x30\x58\x65\x6a\x75\x35\x48\x73\x35\x51\x43\x74\x49\x74\x59\x37\x4f\x41\x46\x73\x48\x2f\x67\x51\x4b\x43\x41\x51\x42\x42\x43\x54\x43\x39\x55\x58\x4e\x6a\x4f\x39\x77\x5a\x70\x59\x62\x6f\x0a\x44\x64\x6f\x41\x6b\x53\x6e\x41\x45\x4c\x77\x48\x4a\x6f\x6d\x32\x78\x67\x53\x2b\x65\x30\x34\x34\x48\x31\x52\x6b\x76\x36\x75\x37\x5a\x4f\x32\x49\x31\x63\x4a\x54\x48\x65\x37\x66\x4b\x43\x68\x64\x6b\x59\x4e\x7a\x4c\x57\x32\x6c\x6d\x35\x30\x51\x44\x2b\x31\x66\x0a\x66\x52\x34\x66\x4a\x39\x47\x31\x43\x77\x6c\x51\x45\x70\x48\x36\x7a\x72\x51\x71\x37\x42\x62\x6e\x77\x62\x68\x34\x49\x4f\x58\x72\x4d\x64\x6f\x71\x47\x62\x6f\x6a\x74\x71\x46\x6c\x6a\x5a\x73\x39\x71\x44\x4e\x67\x6f\x66\x78\x4b\x55\x41\x42\x49\x69\x55\x33\x4b\x0a\x70\x64\x45\x70\x59\x70\x4e\x44\x53\x52\x4f\x5a\x79\x55\x4c\x64\x62\x38\x6c\x39\x74\x75\x75\x35\x4a\x52\x65\x77\x63\x75\x68\x67\x51\x67\x65\x6c\x32\x6b\x6b\x32\x72\x4f\x7a\x4d\x38\x42\x70\x2b\x75\x70\x37\x4b\x75\x59\x42\x6d\x6e\x79\x51\x4a\x43\x65\x55\x6f\x0a\x65\x30\x35\x79\x2f\x73\x66\x38\x31\x73\x76\x2b\x67\x47\x72\x70\x42\x7a\x4c\x74\x77\x69\x45\x7a\x7a\x78\x46\x32\x63\x2f\x46\x71\x6e\x6e\x47\x77\x78\x46\x76\x33\x5a\x33\x42\x72\x6b\x56\x6d\x35\x65\x62\x67\x6f\x65\x5a\x2f\x6c\x30\x41\x58\x6b\x7a\x4d\x6c\x43\x0a\x33\x77\x58\x6f\x50\x62\x46\x4a\x73\x78\x46\x5a\x61\x47\x4a\x37\x7a\x50\x32\x32\x64\x42\x44\x47\x67\x4e\x34\x6f\x56\x4d\x6e\x43\x77\x73\x70\x33\x41\x4b\x55\x4e\x38\x75\x38\x64\x38\x6d\x6a\x55\x6c\x44\x64\x69\x39\x5a\x48\x35\x54\x43\x36\x58\x46\x7a\x42\x54\x0a\x35\x7a\x41\x46\x41\x6f\x49\x42\x41\x46\x43\x6a\x6c\x58\x6d\x63\x30\x4d\x66\x56\x30\x39\x36\x69\x42\x59\x59\x79\x58\x30\x61\x4e\x54\x70\x2f\x6e\x51\x34\x79\x4c\x52\x63\x6e\x37\x49\x66\x6d\x73\x68\x59\x44\x68\x47\x2b\x76\x6f\x6e\x66\x6b\x4b\x46\x4d\x74\x6f\x0a\x31\x38\x31\x39\x67\x48\x61\x61\x47\x78\x57\x4d\x74\x46\x55\x46\x66\x2b\x57\x4f\x4d\x59\x36\x59\x4b\x39\x42\x77\x37\x57\x59\x47\x53\x4b\x48\x4a\x57\x58\x4c\x71\x6d\x42\x4e\x31\x43\x55\x68\x37\x48\x56\x71\x30\x76\x4d\x74\x79\x58\x36\x76\x74\x56\x2f\x51\x51\x0a\x55\x55\x67\x37\x6d\x6f\x76\x61\x75\x30\x42\x75\x75\x48\x70\x38\x6e\x70\x45\x44\x51\x68\x54\x55\x4f\x55\x4e\x47\x30\x4f\x4e\x2b\x34\x43\x62\x59\x5a\x33\x64\x4b\x62\x57\x74\x41\x5a\x56\x65\x48\x4e\x61\x63\x47\x34\x38\x53\x31\x71\x77\x45\x5a\x50\x4c\x31\x75\x0a\x55\x69\x55\x54\x73\x74\x54\x4e\x71\x39\x59\x53\x67\x6b\x49\x2b\x59\x46\x67\x77\x65\x43\x41\x47\x47\x50\x63\x6f\x75\x52\x42\x31\x46\x43\x64\x71\x44\x7a\x50\x48\x6b\x63\x76\x56\x2f\x58\x38\x65\x66\x5a\x51\x55\x49\x54\x73\x53\x47\x4d\x2b\x77\x6e\x58\x57\x30\x0a\x47\x6a\x38\x65\x33\x38\x5a\x63\x4a\x76\x57\x49\x30\x34\x6d\x50\x6f\x44\x30\x50\x39\x57\x61\x4c\x68\x2f\x53\x34\x34\x70\x2b\x52\x45\x6c\x6a\x55\x39\x74\x47\x4a\x6c\x58\x7a\x71\x4c\x32\x6d\x4e\x6d\x6c\x63\x79\x66\x56\x79\x44\x7a\x72\x68\x2b\x67\x41\x4a\x50\x0a\x7a\x59\x71\x36\x75\x41\x58\x63\x7a\x4e\x77\x66\x2f\x55\x46\x2f\x6a\x36\x6f\x43\x4a\x38\x32\x61\x56\x32\x7a\x30\x56\x77\x45\x43\x67\x67\x45\x42\x41\x4a\x67\x61\x32\x4b\x62\x31\x6f\x76\x4e\x4c\x4d\x51\x38\x56\x75\x6d\x31\x74\x45\x38\x42\x4f\x6b\x75\x39\x44\x0a\x68\x54\x56\x70\x7a\x2b\x73\x46\x75\x67\x5a\x75\x59\x35\x44\x7a\x6c\x39\x73\x50\x47\x54\x33\x55\x58\x47\x6a\x5a\x5a\x38\x70\x47\x45\x47\x70\x4c\x50\x51\x4b\x32\x59\x50\x35\x79\x63\x2b\x65\x68\x43\x6e\x67\x33\x65\x50\x74\x58\x56\x30\x73\x50\x55\x52\x35\x57\x0a\x31\x37\x6a\x69\x62\x4d\x32\x4a\x59\x46\x6b\x52\x39\x52\x4a\x43\x43\x41\x38\x63\x2b\x2b\x32\x4f\x77\x78\x5a\x54\x56\x4a\x63\x53\x44\x55\x6b\x48\x47\x37\x63\x70\x6c\x35\x51\x4c\x62\x7a\x4f\x4a\x41\x57\x47\x37\x38\x45\x7a\x69\x69\x4e\x71\x5a\x45\x79\x43\x30\x0a\x52\x73\x59\x4b\x50\x61\x64\x36\x65\x30\x65\x6e\x59\x74\x50\x76\x59\x6c\x6e\x43\x59\x72\x37\x37\x59\x34\x4b\x4e\x36\x6d\x55\x35\x4f\x74\x33\x4f\x7a\x71\x47\x65\x50\x34\x2b\x2b\x64\x55\x4d\x7a\x6e\x4c\x32\x31\x76\x62\x32\x77\x6d\x4f\x36\x78\x48\x4c\x52\x30\x0a\x4d\x77\x61\x36\x77\x37\x66\x49\x58\x59\x4e\x41\x4f\x75\x43\x6b\x30\x47\x4d\x65\x59\x7a\x53\x48\x5a\x54\x37\x5a\x59\x79\x74\x55\x55\x68\x69\x69\x54\x56\x62\x48\x5a\x75\x61\x7a\x4a\x52\x36\x4a\x79\x52\x37\x79\x66\x33\x64\x6e\x6a\x6b\x50\x66\x52\x56\x48\x65\x0a\x71\x4c\x73\x4b\x49\x56\x4e\x62\x79\x57\x61\x79\x6f\x70\x36\x31\x66\x4e\x5a\x57\x66\x4e\x34\x46\x67\x4c\x4d\x6d\x55\x44\x7a\x43\x4e\x4e\x44\x7a\x39\x73\x46\x58\x30\x39\x4a\x6d\x31\x51\x61\x45\x6f\x2b\x44\x75\x6e\x37\x44\x77\x38\x53\x73\x3d\x0a\x2d\x2d\x2d\x2d\x2d\x45\x4e\x44\x20\x52\x53\x41\x20\x50\x52\x49\x56\x41\x54\x45\x20\x4b\x45\x59\x2d\x2d\x2d\x2d\x2d\x0a"

func assets_tls_snakeoil_key() ([]byte, error) {
	return bindata_read(
		_assets_tls_snakeoil_key,
		"assets/tls/snakeoil.key",
	)
}

var _assets_tls_snakeoil_crt = "\x2d\x2d\x2d\x2d\x2d\x42\x45\x47\x49\x4e\x20\x43\x45\x52\x54\x49\x46\x49\x43\x41\x54\x45\x2d\x2d\x2d\x2d\x2d\x0a\x4d\x49\x49\x46\x44\x44\x43\x43\x41\x76\x51\x43\x41\x51\x45\x77\x44\x51\x59\x4a\x4b\x6f\x5a\x49\x68\x76\x63\x4e\x41\x51\x45\x46\x42\x51\x41\x77\x54\x44\x45\x4c\x4d\x41\x6b\x47\x41\x31\x55\x45\x42\x68\x4d\x43\x56\x56\x4d\x78\x45\x7a\x41\x52\x42\x67\x4e\x56\x0a\x42\x41\x67\x54\x43\x6b\x4e\x68\x62\x47\x6c\x6d\x62\x33\x4a\x75\x61\x57\x45\x78\x45\x6a\x41\x51\x42\x67\x4e\x56\x42\x41\x6f\x54\x43\x57\x35\x6e\x63\x6d\x39\x72\x4c\x6d\x4e\x76\x62\x54\x45\x55\x4d\x42\x49\x47\x41\x31\x55\x45\x41\x78\x51\x4c\x4b\x69\x35\x75\x0a\x5a\x33\x4a\x76\x61\x79\x35\x6a\x62\x32\x30\x77\x48\x68\x63\x4e\x4d\x54\x4d\x77\x4e\x6a\x41\x7a\x4d\x44\x51\x79\x4d\x7a\x41\x77\x57\x68\x63\x4e\x4d\x6a\x4d\x77\x4e\x6a\x41\x78\x4d\x44\x51\x79\x4d\x7a\x41\x77\x57\x6a\x42\x4d\x4d\x51\x73\x77\x43\x51\x59\x44\x0a\x56\x51\x51\x47\x45\x77\x4a\x56\x55\x7a\x45\x54\x4d\x42\x45\x47\x41\x31\x55\x45\x43\x42\x4d\x4b\x51\x32\x46\x73\x61\x57\x5a\x76\x63\x6d\x35\x70\x59\x54\x45\x53\x4d\x42\x41\x47\x41\x31\x55\x45\x43\x68\x4d\x4a\x62\x6d\x64\x79\x62\x32\x73\x75\x59\x32\x39\x74\x0a\x4d\x52\x51\x77\x45\x67\x59\x44\x56\x51\x51\x44\x46\x41\x73\x71\x4c\x6d\x35\x6e\x63\x6d\x39\x72\x4c\x6d\x4e\x76\x62\x54\x43\x43\x41\x69\x49\x77\x44\x51\x59\x4a\x4b\x6f\x5a\x49\x68\x76\x63\x4e\x41\x51\x45\x42\x42\x51\x41\x44\x67\x67\x49\x50\x41\x44\x43\x43\x0a\x41\x67\x6f\x43\x67\x67\x49\x42\x41\x4e\x6f\x75\x58\x43\x36\x77\x4a\x61\x63\x50\x49\x46\x59\x45\x69\x4b\x49\x47\x47\x68\x6c\x69\x53\x68\x78\x73\x79\x31\x38\x31\x4c\x4a\x30\x36\x34\x5a\x54\x44\x33\x41\x52\x4d\x54\x69\x39\x4f\x4e\x4f\x73\x52\x6b\x48\x58\x54\x0a\x6f\x4d\x30\x43\x56\x31\x5a\x4e\x44\x6e\x69\x50\x70\x33\x4a\x49\x4c\x69\x47\x4c\x35\x32\x46\x50\x70\x30\x68\x64\x4a\x65\x67\x51\x2b\x70\x7a\x7a\x2f\x67\x44\x6a\x58\x30\x68\x39\x4d\x43\x57\x65\x78\x4d\x34\x73\x71\x4c\x70\x59\x30\x54\x66\x45\x71\x6b\x39\x6c\x0a\x48\x57\x4b\x50\x4a\x67\x2b\x4f\x44\x4b\x38\x77\x61\x31\x2f\x77\x75\x51\x4f\x4b\x2b\x2b\x6d\x2f\x58\x59\x33\x77\x63\x62\x5a\x67\x4c\x6a\x32\x33\x45\x48\x63\x47\x79\x62\x6c\x2f\x78\x47\x4a\x6b\x61\x65\x61\x71\x47\x4e\x76\x72\x4a\x6f\x2f\x34\x52\x67\x53\x4c\x0a\x2f\x6a\x54\x70\x70\x30\x64\x44\x57\x4b\x55\x6d\x4b\x42\x70\x54\x48\x34\x50\x75\x44\x36\x62\x6b\x6c\x38\x71\x4b\x51\x41\x50\x31\x65\x75\x6c\x36\x5a\x53\x2b\x47\x72\x2f\x6c\x6c\x46\x59\x4a\x57\x71\x77\x67\x64\x33\x53\x44\x33\x57\x68\x2f\x36\x31\x59\x34\x6e\x0a\x36\x73\x44\x75\x73\x65\x6d\x32\x4e\x51\x6c\x74\x6e\x6c\x46\x77\x65\x35\x56\x53\x31\x4c\x66\x6f\x70\x56\x74\x4f\x33\x65\x32\x4e\x35\x66\x76\x65\x45\x55\x64\x62\x67\x32\x59\x44\x73\x37\x65\x76\x4e\x59\x56\x4b\x2b\x31\x69\x30\x6b\x6a\x6f\x30\x64\x44\x33\x44\x0a\x74\x62\x6d\x68\x65\x31\x33\x4f\x51\x2f\x62\x61\x75\x66\x78\x63\x48\x63\x31\x78\x6a\x74\x47\x73\x4a\x35\x37\x77\x73\x38\x6e\x65\x50\x4c\x34\x65\x6f\x2b\x76\x46\x75\x44\x35\x51\x6a\x63\x39\x45\x53\x66\x7a\x75\x6e\x53\x65\x6d\x30\x68\x43\x35\x41\x4a\x67\x46\x0a\x72\x73\x57\x7a\x49\x59\x73\x72\x37\x55\x73\x69\x75\x4b\x47\x34\x67\x51\x64\x44\x31\x37\x4c\x2b\x44\x6f\x67\x56\x43\x71\x52\x63\x64\x59\x58\x75\x67\x72\x50\x45\x4a\x4f\x4e\x69\x38\x59\x47\x33\x65\x52\x52\x50\x2f\x51\x39\x6b\x6c\x58\x34\x50\x57\x79\x37\x4f\x0a\x5a\x6a\x37\x4b\x58\x2f\x78\x42\x72\x76\x53\x35\x47\x65\x44\x4e\x71\x72\x44\x56\x73\x50\x4c\x30\x57\x32\x34\x46\x4a\x2f\x2f\x53\x75\x73\x37\x42\x53\x46\x30\x66\x63\x35\x44\x53\x59\x71\x74\x44\x34\x5a\x67\x66\x31\x47\x52\x69\x30\x2b\x78\x53\x61\x33\x6d\x39\x0a\x41\x67\x4a\x77\x57\x4f\x48\x6d\x36\x69\x43\x30\x6d\x48\x77\x64\x64\x6f\x7a\x50\x4c\x64\x2b\x43\x50\x2f\x4d\x37\x50\x59\x7a\x50\x52\x44\x69\x57\x59\x2f\x65\x68\x6f\x43\x46\x69\x71\x52\x41\x67\x6c\x59\x4f\x30\x6d\x68\x4b\x79\x4a\x76\x6f\x54\x4f\x6a\x71\x36\x0a\x54\x42\x59\x59\x70\x2f\x38\x4c\x63\x62\x71\x66\x71\x55\x6b\x49\x36\x44\x53\x6e\x5a\x32\x66\x4e\x64\x58\x4b\x44\x79\x69\x72\x4d\x53\x4f\x54\x34\x4b\x31\x75\x67\x31\x4d\x48\x61\x71\x38\x6f\x62\x75\x63\x4c\x71\x78\x2b\x35\x48\x69\x65\x35\x64\x36\x43\x4b\x52\x0a\x67\x6e\x4e\x31\x35\x73\x4a\x77\x47\x4b\x49\x79\x32\x37\x66\x67\x36\x4c\x41\x7a\x69\x31\x47\x57\x75\x43\x7a\x5a\x31\x6b\x48\x71\x6a\x75\x51\x68\x33\x6f\x56\x66\x75\x42\x4b\x6e\x46\x66\x73\x64\x4c\x55\x6e\x46\x41\x67\x4d\x42\x41\x41\x45\x77\x44\x51\x59\x4a\x0a\x4b\x6f\x5a\x49\x68\x76\x63\x4e\x41\x51\x45\x46\x42\x51\x41\x44\x67\x67\x49\x42\x41\x44\x61\x65\x6c\x32\x2b\x53\x52\x38\x4c\x43\x35\x78\x4e\x43\x42\x63\x75\x34\x65\x71\x6d\x69\x6c\x33\x30\x57\x68\x62\x37\x71\x71\x72\x31\x36\x45\x48\x55\x2f\x4d\x53\x51\x62\x0a\x6e\x6d\x46\x4d\x78\x42\x71\x54\x61\x32\x42\x38\x52\x44\x5a\x49\x49\x6b\x62\x37\x4c\x75\x6b\x48\x38\x72\x73\x41\x64\x55\x38\x42\x7a\x6b\x63\x32\x79\x52\x64\x6a\x6f\x41\x66\x45\x4d\x4d\x63\x4a\x41\x2f\x66\x4d\x70\x77\x75\x61\x58\x49\x35\x63\x75\x61\x4b\x56\x0a\x69\x64\x5a\x46\x70\x4e\x55\x79\x52\x2b\x4b\x35\x55\x47\x2f\x43\x6e\x74\x63\x43\x76\x77\x7a\x5a\x70\x34\x2f\x2f\x67\x2b\x4c\x56\x4b\x39\x71\x50\x44\x5a\x32\x42\x4a\x6d\x41\x2f\x50\x4d\x52\x36\x4f\x70\x68\x52\x77\x52\x77\x47\x2b\x72\x75\x53\x4c\x55\x43\x69\x0a\x79\x77\x67\x46\x46\x68\x4e\x6c\x50\x4d\x70\x50\x5a\x34\x39\x76\x73\x46\x6d\x2f\x51\x30\x41\x34\x4a\x6d\x4c\x70\x5a\x58\x61\x41\x52\x74\x35\x33\x7a\x4e\x62\x4e\x69\x48\x54\x31\x46\x67\x54\x50\x2f\x39\x4c\x31\x48\x49\x70\x6b\x62\x66\x6f\x46\x51\x54\x36\x52\x0a\x70\x42\x2f\x56\x59\x65\x38\x4f\x2b\x47\x6d\x72\x77\x61\x4c\x33\x6c\x2b\x4c\x38\x61\x4f\x30\x37\x4a\x52\x4d\x2b\x75\x30\x4f\x4b\x4e\x50\x78\x4f\x67\x78\x57\x67\x45\x37\x55\x6e\x67\x69\x57\x75\x6c\x58\x70\x6e\x49\x48\x41\x58\x59\x78\x54\x56\x4d\x51\x38\x56\x0a\x31\x49\x59\x75\x74\x76\x39\x62\x59\x4a\x2f\x2b\x54\x41\x79\x39\x4b\x78\x7a\x63\x34\x51\x33\x33\x4b\x2b\x35\x71\x76\x53\x35\x38\x47\x43\x41\x6a\x4f\x46\x31\x30\x4e\x6e\x48\x6b\x69\x54\x51\x50\x62\x68\x55\x57\x49\x53\x6f\x6a\x32\x31\x65\x5a\x4a\x31\x68\x62\x0a\x5a\x34\x42\x62\x43\x35\x7a\x2f\x59\x38\x7a\x4e\x62\x4d\x61\x6f\x30\x41\x43\x46\x30\x51\x6d\x6c\x6e\x58\x6c\x74\x30\x59\x72\x6b\x6a\x6b\x63\x5a\x51\x75\x71\x65\x72\x7a\x70\x43\x36\x59\x4d\x77\x33\x32\x76\x73\x65\x6d\x52\x45\x63\x35\x35\x58\x39\x55\x51\x4d\x0a\x6d\x56\x62\x6e\x44\x61\x68\x34\x78\x79\x50\x4d\x2b\x79\x4e\x49\x67\x2f\x75\x4b\x48\x56\x54\x59\x4a\x36\x2b\x34\x54\x74\x74\x57\x45\x70\x31\x4a\x47\x52\x61\x50\x32\x45\x4f\x59\x72\x30\x79\x47\x65\x34\x58\x46\x46\x72\x4b\x49\x50\x48\x32\x44\x6c\x41\x70\x4d\x0a\x6e\x4e\x66\x71\x52\x2b\x4d\x66\x78\x39\x57\x5a\x53\x33\x6e\x36\x46\x75\x57\x50\x47\x57\x55\x44\x42\x79\x65\x32\x66\x64\x78\x4f\x6a\x57\x39\x4a\x77\x63\x2f\x2b\x4a\x44\x52\x33\x42\x73\x53\x36\x35\x4c\x6c\x7a\x50\x72\x73\x36\x43\x38\x54\x63\x4e\x6e\x6e\x44\x0a\x45\x64\x57\x4a\x34\x6b\x6c\x59\x7a\x57\x6b\x75\x53\x64\x55\x69\x56\x31\x45\x58\x62\x73\x42\x31\x73\x53\x49\x4b\x6d\x55\x75\x64\x32\x66\x34\x76\x4a\x75\x4f\x71\x6c\x42\x73\x67\x53\x38\x2f\x6d\x54\x78\x6a\x6b\x31\x32\x33\x75\x58\x61\x4e\x39\x7a\x61\x4e\x36\x0a\x41\x39\x63\x4d\x57\x51\x49\x32\x4d\x62\x6f\x62\x4b\x30\x48\x45\x72\x6b\x6b\x30\x51\x79\x4e\x54\x74\x56\x54\x4b\x75\x6d\x45\x45\x6b\x6f\x2f\x63\x32\x6b\x74\x73\x6e\x38\x6c\x72\x64\x4a\x63\x4b\x71\x43\x65\x51\x39\x45\x4b\x77\x70\x36\x72\x30\x4c\x6d\x74\x70\x0a\x2d\x2d\x2d\x2d\x2d\x45\x4e\x44\x20\x43\x45\x52\x54\x49\x46\x49\x43\x41\x54\x45\x2d\x2d\x2d\x2d\x2d\x0a"

func assets_tls_snakeoil_crt() ([]byte, error) {
	return bindata_read(
		_assets_tls_snakeoil_crt,
		"assets/tls/snakeoil.crt",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"assets/tls/trusted.root.crt": assets_tls_trusted_root_crt,
	"assets/tls/snakeoil.root.crt": assets_tls_snakeoil_root_crt,
	"assets/tls/snakeoil.key": assets_tls_snakeoil_key,
	"assets/tls/snakeoil.crt": assets_tls_snakeoil_crt,

}
