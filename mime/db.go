package mime

var trueVal = true
var falseVal = false

// source https://github.com/jshttp/mime-db
var mimeTypes = map[string]*Spec{
	"application/rtf": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rtf"},
	},
	"application/vnd.ms-word.template.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"dotm"},
	},
	"application/vnd.ufdl": &Spec{
		Source:     "iana",
		Extensions: []string{"ufd", "ufdl"},
	},
	"audio/x-flac": &Spec{
		Source:     "apache",
		Extensions: []string{"flac"},
	},
	"application/vnd.crick.clicker": &Spec{
		Source:     "iana",
		Extensions: []string{"clkx"},
	},
	"application/vnd.dpgraph": &Spec{
		Source:     "iana",
		Extensions: []string{"dpg"},
	},
	"application/swid+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"swidtag"},
	},
	"image/heic-sequence": &Spec{
		Source:     "iana",
		Extensions: []string{"heics"},
	},
	"video/x-ms-asf": &Spec{
		Source:     "apache",
		Extensions: []string{"asf", "asx"},
	},
	"application/emma+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"emma"},
	},
	"application/vnd.recordare.musicxml": &Spec{
		Source:     "iana",
		Extensions: []string{"mxl"},
	},
	"font/woff": &Spec{
		Source:     "iana",
		Extensions: []string{"woff"},
	},
	"application/gzip": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"gz"},
	},
	"application/vnd.nokia.n-gage.symbian.install": &Spec{
		Source:     "iana",
		Extensions: []string{"n-gage"},
	},
	"application/zip": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"zip"},
	},
	"application/lost+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"lostxml"},
	},
	"application/vnd.lotus-wordpro": &Spec{
		Source:     "iana",
		Extensions: []string{"lwp"},
	},
	"image/vnd.wap.wbmp": &Spec{
		Source:     "iana",
		Extensions: []string{"wbmp"},
	},
	"text/vnd.fmi.flexstor": &Spec{
		Source:     "iana",
		Extensions: []string{"flx"},
	},
	"application/vnd.ms-xpsdocument": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"xps"},
	},
	"application/mac-binhex40": &Spec{
		Source:     "iana",
		Extensions: []string{"hqx"},
	},
	"image/jxra": &Spec{
		Source:     "iana",
		Extensions: []string{"jxra"},
	},
	"text/x-scss": &Spec{
		Extensions: []string{"scss"},
	},
	"application/ld+json": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"jsonld"},
	},
	"video/mp2t": &Spec{
		Source:     "iana",
		Extensions: []string{"ts"},
	},
	"image/vnd.fujixerox.edmics-mmr": &Spec{
		Source:     "iana",
		Extensions: []string{"mmr"},
	},
	"model/vnd.dwf": &Spec{
		Source:     "iana",
		Extensions: []string{"dwf"},
	},
	"application/dssc+der": &Spec{
		Source:     "iana",
		Extensions: []string{"dssc"},
	},
	"application/pskc+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"pskcxml"},
	},
	"application/vnd.cluetrust.cartomobile-config-pkg": &Spec{
		Source:     "iana",
		Extensions: []string{"c11amz"},
	},
	"application/vnd.recordare.musicxml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"musicxml"},
	},
	"application/x-msclip": &Spec{
		Source:     "apache",
		Extensions: []string{"clp"},
	},
	"application/vnd.apple.mpegurl": &Spec{
		Source:     "iana",
		Extensions: []string{"m3u8"},
	},
	"application/vnd.google-earth.kmz": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"kmz"},
	},
	"application/vnd.yamaha.hv-voice": &Spec{
		Source:     "iana",
		Extensions: []string{"hvp"},
	},
	"image/vnd.dece.graphic": &Spec{
		Source:     "iana",
		Extensions: []string{"uvi", "uvvi", "uvg", "uvvg"},
	},
	"application/atsc-rsat+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rsat"},
	},
	"application/xhtml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xhtml", "xht"},
	},
	"image/hej2k": &Spec{
		Source:     "iana",
		Extensions: []string{"hej2"},
	},
	"application/vnd.groove-identity-message": &Spec{
		Source:     "iana",
		Extensions: []string{"gim"},
	},
	"application/vnd.oasis.opendocument.formula": &Spec{
		Source:     "iana",
		Extensions: []string{"odf"},
	},
	"video/jpm": &Spec{
		Source:     "apache",
		Extensions: []string{"jpm", "jpgm"},
	},
	"application/pgp-encrypted": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"pgp"},
	},
	"application/oebps-package+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"opf"},
	},
	"application/vnd.shana.informed.formdata": &Spec{
		Source:     "iana",
		Extensions: []string{"ifm"},
	},
	"application/fdt+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"fdt"},
	},
	"application/mac-compactpro": &Spec{
		Source:     "apache",
		Extensions: []string{"cpt"},
	},
	"application/resource-lists-diff+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rld"},
	},
	"application/x-futuresplash": &Spec{
		Source:     "apache",
		Extensions: []string{"spl"},
	},
	"model/vnd.parasolid.transmit.text": &Spec{
		Source:     "iana",
		Extensions: []string{"x_t"},
	},
	"application/vnd.vsf": &Spec{
		Source:     "iana",
		Extensions: []string{"vsf"},
	},
	"application/vnd.yamaha.smaf-phrase": &Spec{
		Source:     "iana",
		Extensions: []string{"spf"},
	},
	"audio/x-pn-realaudio": &Spec{
		Source:     "apache",
		Extensions: []string{"ram", "ra"},
	},
	"text/csv": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"csv"},
	},
	"text/jsx": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"jsx"},
	},
	"application/vnd.adobe.xdp+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xdp"},
	},
	"application/vnd.ms-excel.sheet.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"xlsm"},
	},
	"application/inkml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ink", "inkml"},
	},
	"application/x-conference": &Spec{
		Source:     "apache",
		Extensions: []string{"nsc"},
	},
	"text/x-org": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"org"},
	},
	"application/vnd.fujitsu.oasys3": &Spec{
		Source:     "iana",
		Extensions: []string{"oa3"},
	},
	"application/oda": &Spec{
		Source:     "iana",
		Extensions: []string{"oda"},
	},
	"chemical/x-xyz": &Spec{
		Source:     "apache",
		Extensions: []string{"xyz"},
	},
	"application/vnd.ms-excel.template.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"xltm"},
	},
	"text/mdx": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"mdx"},
	},
	"application/x-bittorrent": &Spec{
		Source:     "apache",
		Extensions: []string{"torrent"},
	},
	"application/vnd.airzip.filesecure.azf": &Spec{
		Source:     "iana",
		Extensions: []string{"azf"},
	},
	"application/vnd.nokia.n-gage.data": &Spec{
		Source:     "iana",
		Extensions: []string{"ngdat"},
	},
	"image/x-xbitmap": &Spec{
		Source:     "apache",
		Extensions: []string{"xbm"},
	},
	"video/mj2": &Spec{
		Source:     "iana",
		Extensions: []string{"mj2", "mjp2"},
	},
	"application/x-tcl": &Spec{
		Source:     "apache",
		Extensions: []string{"tcl", "tk"},
	},
	"image/jpx": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"jpx", "jpf"},
	},
	"application/vnd.pmi.widget": &Spec{
		Source:     "iana",
		Extensions: []string{"wg"},
	},
	"audio/x-aiff": &Spec{
		Source:     "apache",
		Extensions: []string{"aif", "aiff", "aifc"},
	},
	"audio/x-ms-wax": &Spec{
		Source:     "apache",
		Extensions: []string{"wax"},
	},
	"video/vnd.fvt": &Spec{
		Source:     "iana",
		Extensions: []string{"fvt"},
	},
	"application/vnd.epson.ssf": &Spec{
		Source:     "iana",
		Extensions: []string{"ssf"},
	},
	"text/vnd.graphviz": &Spec{
		Source:     "iana",
		Extensions: []string{"gv"},
	},
	"image/bmp": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"bmp"},
	},
	"image/vnd.ms-dds": &Spec{
		Extensions: []string{"dds"},
	},
	"application/x-rar-compressed": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"rar"},
	},
	"application/applixware": &Spec{
		Source:     "apache",
		Extensions: []string{"aw"},
	},
	"application/docbook+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"dbk"},
	},
	"application/vnd.cloanto.rp9": &Spec{
		Source:     "iana",
		Extensions: []string{"rp9"},
	},
	"image/vnd.dwg": &Spec{
		Source:     "iana",
		Extensions: []string{"dwg"},
	},
	"image/webp": &Spec{
		Source:     "apache",
		Extensions: []string{"webp"},
	},
	"application/srgs+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"grxml"},
	},
	"image/aces": &Spec{
		Source:     "iana",
		Extensions: []string{"exr"},
	},
	"application/vnd.oasis.opendocument.text": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"odt"},
	},
	"application/vnd.sun.xml.impress.template": &Spec{
		Source:     "apache",
		Extensions: []string{"sti"},
	},
	"application/vnd.tmobile-livetv": &Spec{
		Source:     "iana",
		Extensions: []string{"tmo"},
	},
	"application/vnd.acucorp": &Spec{
		Source:     "iana",
		Extensions: []string{"atc", "acutc"},
	},
	"application/vnd.palm": &Spec{
		Source:     "iana",
		Extensions: []string{"pdb", "pqa", "oprc"},
	},
	"application/x-font-bdf": &Spec{
		Source:     "apache",
		Extensions: []string{"bdf"},
	},
	"audio/midi": &Spec{
		Source:     "apache",
		Extensions: []string{"mid", "midi", "kar", "rmi"},
	},
	"application/vnd.data-vision.rdz": &Spec{
		Source:     "iana",
		Extensions: []string{"rdz"},
	},
	"application/vnd.rig.cryptonote": &Spec{
		Source:     "iana",
		Extensions: []string{"cryptonote"},
	},
	"application/x-hdf": &Spec{
		Source:     "apache",
		Extensions: []string{"hdf"},
	},
	"application/xproc+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"xpl"},
	},
	"application/vnd.mfmp": &Spec{
		Source:     "iana",
		Extensions: []string{"mfm"},
	},
	"application/vnd.jisp": &Spec{
		Source:     "iana",
		Extensions: []string{"jisp"},
	},
	"application/vnd.sun.xml.calc.template": &Spec{
		Source:     "apache",
		Extensions: []string{"stc"},
	},
	"application/vnd.hp-hpid": &Spec{
		Source:     "iana",
		Extensions: []string{"hpid"},
	},
	"application/x-cpio": &Spec{
		Source:     "apache",
		Extensions: []string{"cpio"},
	},
	"application/vnd.sun.xml.draw.template": &Spec{
		Source:     "apache",
		Extensions: []string{"std"},
	},
	"image/emf": &Spec{
		Source:     "iana",
		Extensions: []string{"emf"},
	},
	"application/atomsvc+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"atomsvc"},
	},
	"audio/mpeg": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"mpga", "mp2", "mp2a", "mp3", "m2a", "m3a"},
	},
	"audio/vnd.ms-playready.media.pya": &Spec{
		Source:     "iana",
		Extensions: []string{"pya"},
	},
	"application/vnd.simtech-mindmapper": &Spec{
		Source:     "iana",
		Extensions: []string{"twd", "twds"},
	},
	"application/x-msmetafile": &Spec{
		Source:     "apache",
		Extensions: []string{"wmf", "wmz", "emf", "emz"},
	},
	"application/vnd.kde.kchart": &Spec{
		Source:     "iana",
		Extensions: []string{"chrt"},
	},
	"audio/silk": &Spec{
		Source:     "apache",
		Extensions: []string{"sil"},
	},
	"application/mets+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mets"},
	},
	"application/vnd.hbci": &Spec{
		Source:     "iana",
		Extensions: []string{"hbci"},
	},
	"application/x-mspublisher": &Spec{
		Source:     "apache",
		Extensions: []string{"pub"},
	},
	"image/jxss": &Spec{
		Source:     "iana",
		Extensions: []string{"jxss"},
	},
	"application/vnd.llamagraphics.life-balance.exchange+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"lbe"},
	},
	"application/x-eva": &Spec{
		Source:     "apache",
		Extensions: []string{"eva"},
	},
	"image/jxr": &Spec{
		Source:     "iana",
		Extensions: []string{"jxr"},
	},
	"application/vnd.medcalcdata": &Spec{
		Source:     "iana",
		Extensions: []string{"mc1"},
	},
	"application/vnd.muvee.style": &Spec{
		Source:     "iana",
		Extensions: []string{"msty"},
	},
	"application/x-nzb": &Spec{
		Source:     "apache",
		Extensions: []string{"nzb"},
	},
	"application/atom+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"atom"},
	},
	"application/json": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"json", "map"},
	},
	"application/vnd.enliven": &Spec{
		Source:     "iana",
		Extensions: []string{"nml"},
	},
	"application/vnd.ecowin.chart": &Spec{
		Source:     "iana",
		Extensions: []string{"mag"},
	},
	"application/vnd.accpac.simply.aso": &Spec{
		Source:     "iana",
		Extensions: []string{"aso"},
	},
	"application/vnd.openxmlformats-officedocument.wordprocessingml.template": &Spec{
		Source:     "iana",
		Extensions: []string{"dotx"},
	},
	"application/x-cfs-compressed": &Spec{
		Source:     "apache",
		Extensions: []string{"cfs"},
	},
	"application/rss+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"rss"},
	},
	"application/vnd.chipnuts.karaoke-mmd": &Spec{
		Source:     "iana",
		Extensions: []string{"mmd"},
	},
	"application/vnd.powerbuilder6": &Spec{
		Source:     "iana",
		Extensions: []string{"pbd"},
	},
	"image/t38": &Spec{
		Source:     "iana",
		Extensions: []string{"t38"},
	},
	"application/x-iso9660-image": &Spec{
		Source:     "apache",
		Extensions: []string{"iso"},
	},
	"video/x-mng": &Spec{
		Source:     "apache",
		Extensions: []string{"mng"},
	},
	"application/vnd.cluetrust.cartomobile-config": &Spec{
		Source:     "iana",
		Extensions: []string{"c11amc"},
	},
	"application/vnd.kde.kword": &Spec{
		Source:     "iana",
		Extensions: []string{"kwd", "kwt"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.slide": &Spec{
		Source:     "iana",
		Extensions: []string{"sldx"},
	},
	"application/x-font-snf": &Spec{
		Source:     "apache",
		Extensions: []string{"snf"},
	},
	"application/xliff+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xlf"},
	},
	"application/rpki-manifest": &Spec{
		Source:     "iana",
		Extensions: []string{"mft"},
	},
	"application/vnd.wqd": &Spec{
		Source:     "iana",
		Extensions: []string{"wqd"},
	},
	"video/vnd.dece.video": &Spec{
		Source:     "iana",
		Extensions: []string{"uvv", "uvvv"},
	},
	"application/metalink4+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"meta4"},
	},
	"image/vnd.fastbidsheet": &Spec{
		Source:     "iana",
		Extensions: []string{"fbs"},
	},
	"application/vnd.balsamiq.bmml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"bmml"},
	},
	"application/vnd.ms-pki.stl": &Spec{
		Source:     "apache",
		Extensions: []string{"stl"},
	},
	"application/x-dgc-compressed": &Spec{
		Source:     "apache",
		Extensions: []string{"dgc"},
	},
	"application/vnd.jam": &Spec{
		Source:     "iana",
		Extensions: []string{"jam"},
	},
	"application/atomcat+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"atomcat"},
	},
	"application/gxf": &Spec{
		Source:     "apache",
		Extensions: []string{"gxf"},
	},
	"application/ipfix": &Spec{
		Source:     "iana",
		Extensions: []string{"ipfix"},
	},
	"text/x-nfo": &Spec{
		Source:     "apache",
		Extensions: []string{"nfo"},
	},
	"application/exi": &Spec{
		Source:     "iana",
		Extensions: []string{"exi"},
	},
	"application/vnd.ibm.rights-management": &Spec{
		Source:     "iana",
		Extensions: []string{"irm"},
	},
	"application/vnd.zul": &Spec{
		Source:     "iana",
		Extensions: []string{"zir", "zirz"},
	},
	"application/vnd.framemaker": &Spec{
		Source:     "iana",
		Extensions: []string{"fm", "frame", "maker", "book"},
	},
	"image/x-freehand": &Spec{
		Source:     "apache",
		Extensions: []string{"fh", "fhc", "fh4", "fh5", "fh7"},
	},
	"model/vnd.valve.source.compiled-map": &Spec{
		Source:     "iana",
		Extensions: []string{"bsp"},
	},
	"application/vnd.amazon.ebook": &Spec{
		Source:     "apache",
		Extensions: []string{"azw"},
	},
	"application/x-msmediaview": &Spec{
		Source:     "apache",
		Extensions: []string{"mvb", "m13", "m14"},
	},
	"application/vnd.dvb.ait": &Spec{
		Source:     "iana",
		Extensions: []string{"ait"},
	},
	"video/x-f4v": &Spec{
		Source:     "apache",
		Extensions: []string{"f4v"},
	},
	"application/pics-rules": &Spec{
		Source:     "apache",
		Extensions: []string{"prf"},
	},
	"application/vnd.is-xpr": &Spec{
		Source:     "iana",
		Extensions: []string{"xpr"},
	},
	"application/x-virtualbox-vbox-extpack": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"vbox-extpack"},
	},
	"model/vnd.gdl": &Spec{
		Source:     "iana",
		Extensions: []string{"gdl"},
	},
	"application/gpx+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"gpx"},
	},
	"audio/x-realaudio": &Spec{
		Source:     "nginx",
		Extensions: []string{"ra"},
	},
	"model/vrml": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"wrl", "vrml"},
	},
	"image/dicom-rle": &Spec{
		Source:     "iana",
		Extensions: []string{"drle"},
	},
	"text/x-processing": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"pde"},
	},
	"application/x-mscardfile": &Spec{
		Source:     "apache",
		Extensions: []string{"crd"},
	},
	"application/x-web-app-manifest+json": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"webapp"},
	},
	"audio/xm": &Spec{
		Source:     "apache",
		Extensions: []string{"xm"},
	},
	"image/prs.btif": &Spec{
		Source:     "iana",
		Extensions: []string{"btif"},
	},
	"text/troff": &Spec{
		Source:     "iana",
		Extensions: []string{"t", "tr", "roff", "man", "me", "ms"},
	},
	"application/x-msschedule": &Spec{
		Source:     "apache",
		Extensions: []string{"scd"},
	},
	"application/xv+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mxml", "xhvml", "xvml", "xvm"},
	},
	"application/vnd.oasis.opendocument.presentation": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"odp"},
	},
	"application/vnd.smart.teacher": &Spec{
		Source:     "iana",
		Extensions: []string{"teacher"},
	},
	"application/vnd.3gpp2.tcap": &Spec{
		Source:     "iana",
		Extensions: []string{"tcap"},
	},
	"text/tab-separated-values": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"tsv"},
	},
	"video/x-ms-wm": &Spec{
		Source:     "apache",
		Extensions: []string{"wm"},
	},
	"application/vnd.pawaafile": &Spec{
		Source:     "iana",
		Extensions: []string{"paw"},
	},
	"image/vnd.dxf": &Spec{
		Source:     "iana",
		Extensions: []string{"dxf"},
	},
	"application/manifest+json": &Spec{
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"webmanifest"},
	},
	"application/vnd.oasis.opendocument.formula-template": &Spec{
		Source:     "iana",
		Extensions: []string{"odft"},
	},
	"model/x3d-vrml": &Spec{
		Source:     "iana",
		Extensions: []string{"x3dv"},
	},
	"video/x-ms-wmx": &Spec{
		Source:     "apache",
		Extensions: []string{"wmx"},
	},
	"application/vnd.visio": &Spec{
		Source:     "iana",
		Extensions: []string{"vsd", "vst", "vss", "vsw"},
	},
	"text/xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xml"},
	},
	"audio/vnd.dece.audio": &Spec{
		Source:     "iana",
		Extensions: []string{"uva", "uvva"},
	},
	"application/vnd.mobius.mqy": &Spec{
		Source:     "iana",
		Extensions: []string{"mqy"},
	},
	"text/vnd.curl.dcurl": &Spec{
		Source:     "apache",
		Extensions: []string{"dcurl"},
	},
	"application/vnd.1000minds.decision-model+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"1km"},
	},
	"application/vnd.audiograph": &Spec{
		Source:     "iana",
		Extensions: []string{"aep"},
	},
	"application/vnd.fdf": &Spec{
		Source:     "iana",
		Extensions: []string{"fdf"},
	},
	"application/vnd.sun.xml.writer.template": &Spec{
		Source:     "apache",
		Extensions: []string{"stw"},
	},
	"application/epub+zip": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"epub"},
	},
	"application/vnd.adobe.formscentral.fcdt": &Spec{
		Source:     "iana",
		Extensions: []string{"fcdt"},
	},
	"application/x-xliff+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"xlf"},
	},
	"application/xcap-att+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xav"},
	},
	"application/vnd.hp-hpgl": &Spec{
		Source:     "iana",
		Extensions: []string{"hpgl"},
	},
	"video/vnd.dece.sd": &Spec{
		Source:     "iana",
		Extensions: []string{"uvs", "uvvs"},
	},
	"application/vnd.crick.clicker.keyboard": &Spec{
		Source:     "iana",
		Extensions: []string{"clkk"},
	},
	"application/vnd.osgi.subsystem": &Spec{
		Source:     "iana",
		Extensions: []string{"esa"},
	},
	"application/vnd.apple.pages": &Spec{
		Source:     "iana",
		Extensions: []string{"pages"},
	},
	"application/set-payment-initiation": &Spec{
		Source:     "iana",
		Extensions: []string{"setpay"},
	},
	"application/yang": &Spec{
		Source:     "iana",
		Extensions: []string{"yang"},
	},
	"application/vnd.insors.igm": &Spec{
		Source:     "iana",
		Extensions: []string{"igm"},
	},
	"application/vnd.micrografx.flo": &Spec{
		Source:     "iana",
		Extensions: []string{"flo"},
	},
	"video/vnd.dvb.file": &Spec{
		Source:     "iana",
		Extensions: []string{"dvb"},
	},
	"application/gml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"gml"},
	},
	"application/vnd.geogebra.file": &Spec{
		Source:     "iana",
		Extensions: []string{"ggb"},
	},
	"application/x-stuffit": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"sit"},
	},
	"application/x-pkcs7-certreqresp": &Spec{
		Source:     "apache",
		Extensions: []string{"p7r"},
	},
	"application/marcxml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mrcx"},
	},
	"application/vnd.intu.qbo": &Spec{
		Source:     "iana",
		Extensions: []string{"qbo"},
	},
	"application/vnd.syncml.dm+wbxml": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Extensions:   []string{"bdm"},
	},
	"audio/vnd.digital-winds": &Spec{
		Source:     "iana",
		Extensions: []string{"eol"},
	},
	"image/ktx": &Spec{
		Source:     "iana",
		Extensions: []string{"ktx"},
	},
	"text/x-component": &Spec{
		Source:     "nginx",
		Extensions: []string{"htc"},
	},
	"application/vnd.oasis.opendocument.image": &Spec{
		Source:     "iana",
		Extensions: []string{"odi"},
	},
	"application/x-ms-wmd": &Spec{
		Source:     "apache",
		Extensions: []string{"wmd"},
	},
	"application/vnd.oasis.opendocument.text-template": &Spec{
		Source:     "iana",
		Extensions: []string{"ott"},
	},
	"application/vnd.ms-powerpoint": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"ppt", "pps", "pot"},
	},
	"application/x-install-instructions": &Spec{
		Source:     "apache",
		Extensions: []string{"install"},
	},
	"application/x-msaccess": &Spec{
		Source:     "apache",
		Extensions: []string{"mdb"},
	},
	"audio/mobile-xmf": &Spec{
		Source:     "iana",
		Extensions: []string{"mxmf"},
	},
	"audio/vnd.rip": &Spec{
		Source:     "iana",
		Extensions: []string{"rip"},
	},
	"application/vnd.svd": &Spec{
		Source:     "iana",
		Extensions: []string{"svd"},
	},
	"application/vnd.clonk.c4group": &Spec{
		Source:     "iana",
		Extensions: []string{"c4g", "c4d", "c4f", "c4p", "c4u"},
	},
	"application/vnd.groove-vcard": &Spec{
		Source:     "iana",
		Extensions: []string{"vcg"},
	},
	"application/vnd.yamaha.openscoreformat.osfpvg+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"osfpvg"},
	},
	"image/tiff-fx": &Spec{
		Source:     "iana",
		Extensions: []string{"tfx"},
	},
	"application/vnd.sun.xml.impress": &Spec{
		Source:     "apache",
		Extensions: []string{"sxi"},
	},
	"font/woff2": &Spec{
		Source:     "iana",
		Extensions: []string{"woff2"},
	},
	"application/vnd.groove-tool-message": &Spec{
		Source:     "iana",
		Extensions: []string{"gtm"},
	},
	"application/vnd.sun.xml.writer.global": &Spec{
		Source:     "apache",
		Extensions: []string{"sxg"},
	},
	"text/n3": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"n3"},
	},
	"application/vnd.android.package-archive": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"apk"},
	},
	"application/vnd.fluxtime.clip": &Spec{
		Source:     "iana",
		Extensions: []string{"ftc"},
	},
	"application/vnd.koan": &Spec{
		Source:     "iana",
		Extensions: []string{"skp", "skd", "skt", "skm"},
	},
	"application/vnd.3m.post-it-notes": &Spec{
		Source:     "iana",
		Extensions: []string{"pwn"},
	},
	"application/vnd.astraea-software.iota": &Spec{
		Source:     "iana",
		Extensions: []string{"iota"},
	},
	"application/vnd.criticaltools.wbs+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"wbs"},
	},
	"application/vnd.oasis.opendocument.image-template": &Spec{
		Source:     "iana",
		Extensions: []string{"oti"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"pptx"},
	},
	"application/vnd.adobe.air-application-installer-package+zip": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"air"},
	},
	"application/pkix-pkipath": &Spec{
		Source:     "iana",
		Extensions: []string{"pkipath"},
	},
	"application/route-apd+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rapd"},
	},
	"application/vnd.curl.car": &Spec{
		Source:     "apache",
		Extensions: []string{"car"},
	},
	"application/x-doom": &Spec{
		Source:     "apache",
		Extensions: []string{"wad"},
	},
	"application/vnd.osgi.dp": &Spec{
		Source:     "iana",
		Extensions: []string{"dp"},
	},
	"application/vnd.dynageo": &Spec{
		Source:     "iana",
		Extensions: []string{"geo"},
	},
	"application/vnd.kde.kspread": &Spec{
		Source:     "iana",
		Extensions: []string{"ksp"},
	},
	"image/tiff": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"tif", "tiff"},
	},
	"text/markdown": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"markdown", "md"},
	},
	"application/ogg": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"ogx"},
	},
	"image/x-xpixmap": &Spec{
		Source:     "apache",
		Extensions: []string{"xpm"},
	},
	"text/vnd.curl": &Spec{
		Source:     "iana",
		Extensions: []string{"curl"},
	},
	"image/gif": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"gif"},
	},
	"application/vnd.intergeo": &Spec{
		Source:     "iana",
		Extensions: []string{"i2g"},
	},
	"model/vnd.gtw": &Spec{
		Source:     "iana",
		Extensions: []string{"gtw"},
	},
	"application/onenote": &Spec{
		Source:     "apache",
		Extensions: []string{"onetoc", "onetoc2", "onetmp", "onepkg"},
	},
	"application/vnd.hal+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"hal"},
	},
	"application/x-xz": &Spec{
		Source:     "apache",
		Extensions: []string{"xz"},
	},
	"text/plain": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"txt", "text", "conf", "def", "list", "log", "in", "ini"},
	},
	"application/x-csh": &Spec{
		Source:     "apache",
		Extensions: []string{"csh"},
	},
	"audio/x-caf": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"caf"},
	},
	"text/coffeescript": &Spec{
		Extensions: []string{"coffee", "litcoffee"},
	},
	"application/vnd.oasis.opendocument.graphics-template": &Spec{
		Source:     "iana",
		Extensions: []string{"otg"},
	},
	"application/x-tads": &Spec{
		Source:     "apache",
		Extensions: []string{"gam"},
	},
	"text/x-suse-ymp": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"ymp"},
	},
	"application/vnd.kenameaapp": &Spec{
		Source:     "iana",
		Extensions: []string{"htke"},
	},
	"application/vnd.las.las+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"lasxml"},
	},
	"image/heic": &Spec{
		Source:     "iana",
		Extensions: []string{"heic"},
	},
	"text/jade": &Spec{
		Extensions: []string{"jade"},
	},
	"application/vnd.crick.clicker.template": &Spec{
		Source:     "iana",
		Extensions: []string{"clkt"},
	},
	"application/vnd.oasis.opendocument.text-master": &Spec{
		Source:     "iana",
		Extensions: []string{"odm"},
	},
	"application/vnd.hp-pcl": &Spec{
		Source:     "iana",
		Extensions: []string{"pcl"},
	},
	"application/mxf": &Spec{
		Source:     "iana",
		Extensions: []string{"mxf"},
	},
	"application/shf+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"shf"},
	},
	"application/vnd.novadigm.ext": &Spec{
		Source:     "iana",
		Extensions: []string{"ext"},
	},
	"application/vnd.quark.quarkxpress": &Spec{
		Source:     "iana",
		Extensions: []string{"qxd", "qxt", "qwd", "qwt", "qxl", "qxb"},
	},
	"application/x-latex": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"latex"},
	},
	"font/otf": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"otf"},
	},
	"application/vnd.lotus-1-2-3": &Spec{
		Source:     "iana",
		Extensions: []string{"123"},
	},
	"application/vnd.openblox.game+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"obgx"},
	},
	"video/vnd.uvvu.mp4": &Spec{
		Source:     "iana",
		Extensions: []string{"uvu", "uvvu"},
	},
	"application/vnd.yamaha.hv-dic": &Spec{
		Source:     "iana",
		Extensions: []string{"hvd"},
	},
	"application/vnd.lotus-freelance": &Spec{
		Source:     "iana",
		Extensions: []string{"pre"},
	},
	"application/winhlp": &Spec{
		Source:     "apache",
		Extensions: []string{"hlp"},
	},
	"text/x-c": &Spec{
		Source:     "apache",
		Extensions: []string{"c", "cc", "cxx", "cpp", "h", "hh", "dic"},
	},
	"application/pkcs8": &Spec{
		Source:     "iana",
		Extensions: []string{"p8"},
	},
	"application/vnd.americandynamics.acc": &Spec{
		Source:     "iana",
		Extensions: []string{"acc"},
	},
	"application/x-gnumeric": &Spec{
		Source:     "apache",
		Extensions: []string{"gnumeric"},
	},
	"application/x-mie": &Spec{
		Source:     "apache",
		Extensions: []string{"mie"},
	},
	"audio/x-wav": &Spec{
		Source:     "apache",
		Extensions: []string{"wav"},
	},
	"model/x3d+vrml": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"x3dv", "x3dvz"},
	},
	"text/mathml": &Spec{
		Source:     "nginx",
		Extensions: []string{"mml"},
	},
	"application/javascript": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"js", "mjs"},
	},
	"application/patch-ops-error+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xer"},
	},
	"message/global-disposition-notification": &Spec{
		Source:     "iana",
		Extensions: []string{"u8mdn"},
	},
	"model/x3d+fastinfoset": &Spec{
		Source:     "iana",
		Extensions: []string{"x3db"},
	},
	"application/vnd.previewsystems.box": &Spec{
		Source:     "iana",
		Extensions: []string{"box"},
	},
	"video/vnd.vivo": &Spec{
		Source:     "iana",
		Extensions: []string{"viv"},
	},
	"application/vnd.spotfire.dxp": &Spec{
		Source:     "iana",
		Extensions: []string{"dxp"},
	},
	"font/ttf": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ttf"},
	},
	"video/x-ms-wmv": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"wmv"},
	},
	"application/vnd.wap.wbxml": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Extensions:   []string{"wbxml"},
	},
	"video/x-matroska": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"mkv", "mk3d", "mks"},
	},
	"application/vnd.3gpp.pic-bw-large": &Spec{
		Source:     "iana",
		Extensions: []string{"plb"},
	},
	"application/wspolicy+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"wspolicy"},
	},
	"application/vnd.dart": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"dart"},
	},
	"application/vnd.kinar": &Spec{
		Source:     "iana",
		Extensions: []string{"kne", "knp"},
	},
	"application/x-virtualbox-hdd": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"hdd"},
	},
	"image/jxrs": &Spec{
		Source:     "iana",
		Extensions: []string{"jxrs"},
	},
	"application/vnd.novadigm.edx": &Spec{
		Source:     "iana",
		Extensions: []string{"edx"},
	},
	"image/vnd.dvb.subtitle": &Spec{
		Source:     "iana",
		Extensions: []string{"sub"},
	},
	"application/urc-ressheet+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rsheet"},
	},
	"application/vnd.ms-ims": &Spec{
		Source:     "iana",
		Extensions: []string{"ims"},
	},
	"application/x-msmoney": &Spec{
		Source:     "apache",
		Extensions: []string{"mny"},
	},
	"image/jph": &Spec{
		Source:     "iana",
		Extensions: []string{"jph"},
	},
	"application/sru+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"sru"},
	},
	"application/vnd.crick.clicker.wordbank": &Spec{
		Source:     "iana",
		Extensions: []string{"clkw"},
	},
	"application/vnd.kde.kontour": &Spec{
		Source:     "iana",
		Extensions: []string{"kon"},
	},
	"image/prs.pti": &Spec{
		Source:     "iana",
		Extensions: []string{"pti"},
	},
	"application/vnd.musician": &Spec{
		Source:     "iana",
		Extensions: []string{"mus"},
	},
	"application/vnd.noblenet-sealer": &Spec{
		Source:     "iana",
		Extensions: []string{"nns"},
	},
	"text/vnd.dvb.subtitle": &Spec{
		Source:     "iana",
		Extensions: []string{"sub"},
	},
	"video/x-msvideo": &Spec{
		Source:     "apache",
		Extensions: []string{"avi"},
	},
	"application/vnd.irepository.package+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"irp"},
	},
	"model/x3d+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"x3d", "x3dz"},
	},
	"application/vnd.proteus.magazine": &Spec{
		Source:     "iana",
		Extensions: []string{"mgz"},
	},
	"application/resource-lists+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rl"},
	},
	"application/x-bzip2": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"bz2", "boz"},
	},
	"text/x-uuencode": &Spec{
		Source:     "apache",
		Extensions: []string{"uu"},
	},
	"model/vnd.mts": &Spec{
		Source:     "iana",
		Extensions: []string{"mts"},
	},
	"application/hjson": &Spec{
		Extensions: []string{"hjson"},
	},
	"application/ecmascript": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ecma", "es"},
	},
	"text/rtf": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rtf"},
	},
	"application/vnd.geometry-explorer": &Spec{
		Source:     "iana",
		Extensions: []string{"gex", "gre"},
	},
	"application/vnd.geoplan": &Spec{
		Source:     "iana",
		Extensions: []string{"g2w"},
	},
	"application/x-tgif": &Spec{
		Source:     "apache",
		Extensions: []string{"obj"},
	},
	"application/vnd.stardivision.writer-global": &Spec{
		Source:     "apache",
		Extensions: []string{"sgl"},
	},
	"application/x-apple-diskimage": &Spec{
		Source:     "apache",
		Extensions: []string{"dmg"},
	},
	"application/vnd.mcd": &Spec{
		Source:     "iana",
		Extensions: []string{"mcd"},
	},
	"application/vnd.ms-powerpoint.slide.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"sldm"},
	},
	"image/x-portable-pixmap": &Spec{
		Source:     "apache",
		Extensions: []string{"ppm"},
	},
	"application/sieve": &Spec{
		Source:     "iana",
		Extensions: []string{"siv", "sieve"},
	},
	"application/x-cbr": &Spec{
		Source:     "apache",
		Extensions: []string{"cbr", "cba", "cbt", "cbz", "cb7"},
	},
	"audio/x-pn-realaudio-plugin": &Spec{
		Source:     "apache",
		Extensions: []string{"rmp"},
	},
	"application/vnd.mif": &Spec{
		Source:     "iana",
		Extensions: []string{"mif"},
	},
	"message/rfc822": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"eml", "mime"},
	},
	"application/raml+yaml": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"raml"},
	},
	"application/vnd.ctc-posml": &Spec{
		Source:     "iana",
		Extensions: []string{"pml"},
	},
	"application/vnd.yellowriver-custom-menu": &Spec{
		Source:     "iana",
		Extensions: []string{"cmp"},
	},
	"application/vnd.lotus-notes": &Spec{
		Source:     "iana",
		Extensions: []string{"nsf"},
	},
	"application/vnd.oasis.opendocument.graphics": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"odg"},
	},
	"application/vnd.kahootz": &Spec{
		Source:     "iana",
		Extensions: []string{"ktz", "ktr"},
	},
	"application/x-x509-ca-cert": &Spec{
		Source:     "iana",
		Extensions: []string{"der", "crt", "pem"},
	},
	"model/gltf+json": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"gltf"},
	},
	"text/uri-list": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"uri", "uris", "urls"},
	},
	"application/vnd.apple.pkpass": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"pkpass"},
	},
	"application/x-font-linux-psf": &Spec{
		Source:     "apache",
		Extensions: []string{"psf"},
	},
	"application/x-ns-proxy-autoconfig": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"pac"},
	},
	"application/vnd.stardivision.writer": &Spec{
		Source:     "apache",
		Extensions: []string{"sdw", "vor"},
	},
	"application/vnd.xara": &Spec{
		Source:     "iana",
		Extensions: []string{"xar"},
	},
	"application/x-makeself": &Spec{
		Source:     "nginx",
		Extensions: []string{"run"},
	},
	"application/x-sql": &Spec{
		Source:     "apache",
		Extensions: []string{"sql"},
	},
	"audio/vnd.dts": &Spec{
		Source:     "iana",
		Extensions: []string{"dts"},
	},
	"image/jxsi": &Spec{
		Source:     "iana",
		Extensions: []string{"jxsi"},
	},
	"application/rdf+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rdf", "owl"},
	},
	"application/vnd.oma.dd2+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"dd2"},
	},
	"application/widget": &Spec{
		Source:     "iana",
		Extensions: []string{"wgt"},
	},
	"application/x-mswrite": &Spec{
		Source:     "apache",
		Extensions: []string{"wri"},
	},
	"application/vnd.xfdl": &Spec{
		Source:     "iana",
		Extensions: []string{"xfdl"},
	},
	"application/mathml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mathml"},
	},
	"application/xcap-caps+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xca"},
	},
	"application/java-archive": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"jar", "war", "ear"},
	},
	"application/vnd.publishare-delta-tree": &Spec{
		Source:     "iana",
		Extensions: []string{"qps"},
	},
	"application/vnd.shana.informed.package": &Spec{
		Source:     "iana",
		Extensions: []string{"ipk"},
	},
	"audio/x-m4a": &Spec{
		Source:     "nginx",
		Extensions: []string{"m4a"},
	},
	"image/x-xwindowdump": &Spec{
		Source:     "apache",
		Extensions: []string{"xwd"},
	},
	"application/mrb-consumer+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xdf"},
	},
	"application/vnd.groove-tool-template": &Spec{
		Source:     "iana",
		Extensions: []string{"tpl"},
	},
	"video/ogg": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"ogv"},
	},
	"application/atomdeleted+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"atomdeleted"},
	},
	"application/rsd+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"rsd"},
	},
	"application/vnd.ms-powerpoint.addin.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"ppam"},
	},
	"application/vnd.apple.installer+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mpkg"},
	},
	"application/vnd.adobe.fxp": &Spec{
		Source:     "iana",
		Extensions: []string{"fxp", "fxpl"},
	},
	"application/vnd.fujitsu.oasysgp": &Spec{
		Source:     "iana",
		Extensions: []string{"fg5"},
	},
	"application/vnd.mynfc": &Spec{
		Source:     "iana",
		Extensions: []string{"taglet"},
	},
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"xlsx"},
	},
	"application/vnd.pg.osasli": &Spec{
		Source:     "iana",
		Extensions: []string{"ei6"},
	},
	"application/vnd.bmi": &Spec{
		Source:     "iana",
		Extensions: []string{"bmi"},
	},
	"application/vnd.eszigno3+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"es3", "et3"},
	},
	"application/vnd.ms-wpl": &Spec{
		Source:     "iana",
		Extensions: []string{"wpl"},
	},
	"application/atsc-dwd+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"dwd"},
	},
	"application/font-tdpfr": &Spec{
		Source:     "iana",
		Extensions: []string{"pfr"},
	},
	"application/vnd.lotus-organizer": &Spec{
		Source:     "iana",
		Extensions: []string{"org"},
	},
	"application/vnd.lotus-screencam": &Spec{
		Source:     "iana",
		Extensions: []string{"scm"},
	},
	"application/x-font-ghostscript": &Spec{
		Source:     "apache",
		Extensions: []string{"gsf"},
	},
	"text/vnd.fly": &Spec{
		Source:     "iana",
		Extensions: []string{"fly"},
	},
	"application/smil+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"smi", "smil"},
	},
	"application/vnd.novadigm.edm": &Spec{
		Source:     "iana",
		Extensions: []string{"edm"},
	},
	"application/vnd.oasis.opendocument.spreadsheet": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"ods"},
	},
	"application/x-virtualbox-vhd": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"vhd"},
	},
	"audio/s3m": &Spec{
		Source:     "apache",
		Extensions: []string{"s3m"},
	},
	"application/vnd.ms-powerpoint.template.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"potm"},
	},
	"application/mbox": &Spec{
		Source:     "iana",
		Extensions: []string{"mbox"},
	},
	"application/vnd.3gpp.pic-bw-var": &Spec{
		Source:     "iana",
		Extensions: []string{"pvb"},
	},
	"application/x-virtualbox-ovf": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"ovf"},
	},
	"audio/wav": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"wav"},
	},
	"text/html": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"html", "htm", "shtml"},
	},
	"application/vnd.apple.numbers": &Spec{
		Source:     "iana",
		Extensions: []string{"numbers"},
	},
	"application/vnd.blueice.multipass": &Spec{
		Source:     "iana",
		Extensions: []string{"mpm"},
	},
	"application/x-cocoa": &Spec{
		Source:     "nginx",
		Extensions: []string{"cco"},
	},
	"application/vnd.geonext": &Spec{
		Source:     "iana",
		Extensions: []string{"gxt"},
	},
	"application/vnd.kidspiration": &Spec{
		Source:     "iana",
		Extensions: []string{"kia"},
	},
	"application/x-dtbresource+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"res"},
	},
	"application/cdfx+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"cdfx"},
	},
	"application/vnd.shana.informed.formtemplate": &Spec{
		Source:     "iana",
		Extensions: []string{"itp"},
	},
	"text/x-java-source": &Spec{
		Source:     "apache",
		Extensions: []string{"java"},
	},
	"text/x-setext": &Spec{
		Source:     "apache",
		Extensions: []string{"etx"},
	},
	"application/ssdl+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"ssdl"},
	},
	"application/vnd.epson.msf": &Spec{
		Source:     "iana",
		Extensions: []string{"msf"},
	},
	"text/x-handlebars-template": &Spec{
		Extensions: []string{"hbs"},
	},
	"application/pkix-attr-cert": &Spec{
		Source:     "iana",
		Extensions: []string{"ac"},
	},
	"application/x-dtbook+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"dtb"},
	},
	"application/xcap-diff+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xdf"},
	},
	"application/vnd.dna": &Spec{
		Source:     "iana",
		Extensions: []string{"dna"},
	},
	"application/vnd.geogebra.tool": &Spec{
		Source:     "iana",
		Extensions: []string{"ggt"},
	},
	"application/vnd.ms-powerpoint.presentation.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"pptm"},
	},
	"application/vnd.vcx": &Spec{
		Source:     "iana",
		Extensions: []string{"vcx"},
	},
	"application/x-sv4crc": &Spec{
		Source:     "apache",
		Extensions: []string{"sv4crc"},
	},
	"application/vnd.ipunplugged.rcprofile": &Spec{
		Source:     "iana",
		Extensions: []string{"rcprofile"},
	},
	"application/x-mobipocket-ebook": &Spec{
		Source:     "apache",
		Extensions: []string{"prc", "mobi"},
	},
	"application/x-perl": &Spec{
		Source:     "nginx",
		Extensions: []string{"pl", "pm"},
	},
	"image/svg+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"svg", "svgz"},
	},
	"image/vnd.ms-modi": &Spec{
		Source:     "iana",
		Extensions: []string{"mdi"},
	},
	"application/vnd.semf": &Spec{
		Source:     "iana",
		Extensions: []string{"semf"},
	},
	"application/vnd.openofficeorg.extension": &Spec{
		Source:     "apache",
		Extensions: []string{"oxt"},
	},
	"video/mp4": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"mp4", "mp4v", "mpg4"},
	},
	"application/relax-ng-compact-syntax": &Spec{
		Source:     "iana",
		Extensions: []string{"rnc"},
	},
	"application/vnd.umajin": &Spec{
		Source:     "iana",
		Extensions: []string{"umj"},
	},
	"application/x-msbinder": &Spec{
		Source:     "apache",
		Extensions: []string{"obd"},
	},
	"application/vnd.google-apps.presentation": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"gslides"},
	},
	"image/jphc": &Spec{
		Source:     "iana",
		Extensions: []string{"jhc"},
	},
	"application/prs.cww": &Spec{
		Source:     "iana",
		Extensions: []string{"cww"},
	},
	"application/srgs": &Spec{
		Source:     "iana",
		Extensions: []string{"gram"},
	},
	"application/vnd.openxmlformats-officedocument.spreadsheetml.template": &Spec{
		Source:     "iana",
		Extensions: []string{"xltx"},
	},
	"audio/adpcm": &Spec{
		Source:     "apache",
		Extensions: []string{"adp"},
	},
	"application/vnd.hydrostatix.sof-data": &Spec{
		Source:     "iana",
		Extensions: []string{"sfd-hdstx"},
	},
	"text/vcard": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"vcard"},
	},
	"application/vnd.mobius.dis": &Spec{
		Source:     "iana",
		Extensions: []string{"dis"},
	},
	"application/sdp": &Spec{
		Source:     "iana",
		Extensions: []string{"sdp"},
	},
	"application/vnd.groove-help": &Spec{
		Source:     "iana",
		Extensions: []string{"ghf"},
	},
	"image/jxsc": &Spec{
		Source:     "iana",
		Extensions: []string{"jxsc"},
	},
	"application/bdoc": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"bdoc"},
	},
	"application/node": &Spec{
		Source:     "iana",
		Extensions: []string{"cjs"},
	},
	"application/omdoc+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"omdoc"},
	},
	"model/vnd.usdz+zip": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"usdz"},
	},
	"image/x-portable-anymap": &Spec{
		Source:     "apache",
		Extensions: []string{"pnm"},
	},
	"application/emotionml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"emotionml"},
	},
	"application/vnd.acucobol": &Spec{
		Source:     "iana",
		Extensions: []string{"acu"},
	},
	"application/vnd.gmx": &Spec{
		Source:     "iana",
		Extensions: []string{"gmx"},
	},
	"message/vnd.wfa.wsc": &Spec{
		Source:     "iana",
		Extensions: []string{"wsc"},
	},
	"video/h264": &Spec{
		Source:     "iana",
		Extensions: []string{"h264"},
	},
	"application/vnd.adobe.xfdf": &Spec{
		Source:     "iana",
		Extensions: []string{"xfdf"},
	},
	"application/vnd.uoml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"uoml"},
	},
	"application/x-bcpio": &Spec{
		Source:     "apache",
		Extensions: []string{"bcpio"},
	},
	"application/geo+json": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"geojson"},
	},
	"application/x-pkcs7-certificates": &Spec{
		Source:     "apache",
		Extensions: []string{"p7b", "spc"},
	},
	"application/x-sea": &Spec{
		Source:     "nginx",
		Extensions: []string{"sea"},
	},
	"model/mesh": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"msh", "mesh", "silo"},
	},
	"application/vnd.nokia.radio-presets": &Spec{
		Source:     "iana",
		Extensions: []string{"rpss"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.template": &Spec{
		Source:     "iana",
		Extensions: []string{"potx"},
	},
	"image/x-icon": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"ico"},
	},
	"application/vnd.noblenet-web": &Spec{
		Source:     "iana",
		Extensions: []string{"nnw"},
	},
	"application/vnd.stardivision.calc": &Spec{
		Source:     "apache",
		Extensions: []string{"sdc"},
	},
	"application/vnd.wap.wmlc": &Spec{
		Source:     "iana",
		Extensions: []string{"wmlc"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.slideshow": &Spec{
		Source:     "iana",
		Extensions: []string{"ppsx"},
	},
	"text/richtext": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rtx"},
	},
	"application/vnd.commonspace": &Spec{
		Source:     "iana",
		Extensions: []string{"csp"},
	},
	"application/vnd.rn-realmedia-vbr": &Spec{
		Source:     "apache",
		Extensions: []string{"rmvb"},
	},
	"application/x-silverlight-app": &Spec{
		Source:     "apache",
		Extensions: []string{"xap"},
	},
	"application/vnd.flographit": &Spec{
		Source:     "iana",
		Extensions: []string{"gph"},
	},
	"application/x-virtualbox-vbox": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"vbox"},
	},
	"application/x-shar": &Spec{
		Source:     "apache",
		Extensions: []string{"shar"},
	},
	"text/x-fortran": &Spec{
		Source:     "apache",
		Extensions: []string{"f", "for", "f77", "f90"},
	},
	"application/provenance+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"provx"},
	},
	"text/less": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"less"},
	},
	"application/mp4": &Spec{
		Source:     "iana",
		Extensions: []string{"mp4s", "m4p"},
	},
	"text/slim": &Spec{
		Extensions: []string{"slim", "slm"},
	},
	"application/vnd.amiga.ami": &Spec{
		Source:     "iana",
		Extensions: []string{"ami"},
	},
	"application/vnd.syncml.dm+xml": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"xdm"},
	},
	"application/xcap-error+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xer"},
	},
	"application/vnd.fujitsu.oasys": &Spec{
		Source:     "iana",
		Extensions: []string{"oas"},
	},
	"text/x-opml": &Spec{
		Source:     "apache",
		Extensions: []string{"opml"},
	},
	"application/andrew-inset": &Spec{
		Source:     "iana",
		Extensions: []string{"ez"},
	},
	"application/vnd.ahead.space": &Spec{
		Source:     "iana",
		Extensions: []string{"ahead"},
	},
	"application/x-subrip": &Spec{
		Source:     "apache",
		Extensions: []string{"srt"},
	},
	"application/vnd.fujitsu.oasys2": &Spec{
		Source:     "iana",
		Extensions: []string{"oa2"},
	},
	"application/vnd.dvb.service": &Spec{
		Source:     "iana",
		Extensions: []string{"svc"},
	},
	"text/x-vcalendar": &Spec{
		Source:     "apache",
		Extensions: []string{"vcs"},
	},
	"application/scvp-cv-request": &Spec{
		Source:     "iana",
		Extensions: []string{"scq"},
	},
	"application/vnd.genomatix.tuxedo": &Spec{
		Source:     "iana",
		Extensions: []string{"txd"},
	},
	"application/vnd.pg.format": &Spec{
		Source:     "iana",
		Extensions: []string{"str"},
	},
	"application/vnd.sun.xml.writer": &Spec{
		Source:     "apache",
		Extensions: []string{"sxw"},
	},
	"audio/vnd.nuera.ecelp9600": &Spec{
		Source:     "iana",
		Extensions: []string{"ecelp9600"},
	},
	"image/vnd.net-fpx": &Spec{
		Source:     "iana",
		Extensions: []string{"npx"},
	},
	"application/mads+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mads"},
	},
	"application/vnd.oasis.opendocument.chart-template": &Spec{
		Source:     "iana",
		Extensions: []string{"otc"},
	},
	"model/vnd.collada+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"dae"},
	},
	"application/vnd.claymore": &Spec{
		Source:     "iana",
		Extensions: []string{"cla"},
	},
	"application/x-font-type1": &Spec{
		Source:     "apache",
		Extensions: []string{"pfa", "pfb", "pfm", "afm"},
	},
	"application/vnd.openstreetmap.data+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"osm"},
	},
	"application/x-t3vm-image": &Spec{
		Source:     "apache",
		Extensions: []string{"t3"},
	},
	"text/vnd.wap.wml": &Spec{
		Source:     "iana",
		Extensions: []string{"wml"},
	},
	"application/set-registration-initiation": &Spec{
		Source:     "iana",
		Extensions: []string{"setreg"},
	},
	"application/x-sh": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"sh"},
	},
	"application/vnd.oasis.opendocument.spreadsheet-template": &Spec{
		Source:     "iana",
		Extensions: []string{"ots"},
	},
	"application/vnd.realvnc.bed": &Spec{
		Source:     "iana",
		Extensions: []string{"bed"},
	},
	"text/x-markdown": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"mkd"},
	},
	"text/css": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"css"},
	},
	"application/vnd.olpc-sugar": &Spec{
		Source:     "iana",
		Extensions: []string{"xo"},
	},
	"video/3gpp2": &Spec{
		Source:     "iana",
		Extensions: []string{"3g2"},
	},
	"application/route-usd+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rusd"},
	},
	"application/vnd.cinderella": &Spec{
		Source:     "iana",
		Extensions: []string{"cdy"},
	},
	"application/vnd.immervision-ivp": &Spec{
		Source:     "iana",
		Extensions: []string{"ivp"},
	},
	"application/calendar+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xcs"},
	},
	"text/stylus": &Spec{
		Extensions: []string{"stylus", "styl"},
	},
	"application/rpki-roa": &Spec{
		Source:     "iana",
		Extensions: []string{"roa"},
	},
	"text/vnd.curl.mcurl": &Spec{
		Source:     "apache",
		Extensions: []string{"mcurl"},
	},
	"application/ccxml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ccxml"},
	},
	"application/vnd.webturbo": &Spec{
		Source:     "iana",
		Extensions: []string{"wtb"},
	},
	"application/wsdl+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"wsdl"},
	},
	"text/x-sfv": &Spec{
		Source:     "apache",
		Extensions: []string{"sfv"},
	},
	"application/vnd.ms-excel.addin.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"xlam"},
	},
	"application/vnd.denovo.fcselayout-link": &Spec{
		Source:     "iana",
		Extensions: []string{"fe_launch"},
	},
	"application/vnd.mobius.daf": &Spec{
		Source:     "iana",
		Extensions: []string{"daf"},
	},
	"application/vnd.shana.informed.interchange": &Spec{
		Source:     "iana",
		Extensions: []string{"iif"},
	},
	"application/x-ms-application": &Spec{
		Source:     "apache",
		Extensions: []string{"application"},
	},
	"application/x-freearc": &Spec{
		Source:     "apache",
		Extensions: []string{"arc"},
	},
	"model/x3d+binary": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"x3db", "x3dbz"},
	},
	"application/route-s-tsid+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"sls"},
	},
	"application/vnd.rim.cod": &Spec{
		Source:     "apache",
		Extensions: []string{"cod"},
	},
	"application/vnd.trueapp": &Spec{
		Source:     "iana",
		Extensions: []string{"tra"},
	},
	"video/vnd.dece.pd": &Spec{
		Source:     "iana",
		Extensions: []string{"uvp", "uvvp"},
	},
	"application/vnd.dece.ttml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"uvt", "uvvt"},
	},
	"image/vnd.zbrush.pcx": &Spec{
		Source:     "iana",
		Extensions: []string{"pcx"},
	},
	"video/vnd.mpegurl": &Spec{
		Source:     "iana",
		Extensions: []string{"mxu", "m4u"},
	},
	"image/jpm": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"jpm"},
	},
	"text/calendar": &Spec{
		Source:     "iana",
		Extensions: []string{"ics", "ifb"},
	},
	"application/vnd.sus-calendar": &Spec{
		Source:     "iana",
		Extensions: []string{"sus", "susp"},
	},
	"application/x-netcdf": &Spec{
		Source:     "apache",
		Extensions: []string{"nc", "cdf"},
	},
	"video/h263": &Spec{
		Source:     "iana",
		Extensions: []string{"h263"},
	},
	"application/vnd.curl.pcurl": &Spec{
		Source:     "apache",
		Extensions: []string{"pcurl"},
	},
	"application/vnd.llamagraphics.life-balance.desktop": &Spec{
		Source:     "iana",
		Extensions: []string{"lbd"},
	},
	"application/octet-stream": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"bin", "dms", "lrf", "mar", "so", "dist", "distz", "pkg", "bpk", "dump", "elc", "deploy", "exe", "dll", "deb", "dmg", "iso", "img", "msi", "msp", "msm", "buffer"},
	},
	"application/vnd.cosmocaller": &Spec{
		Source:     "iana",
		Extensions: []string{"cmc"},
	},
	"application/mathematica": &Spec{
		Source:     "iana",
		Extensions: []string{"ma", "nb", "mb"},
	},
	"audio/x-matroska": &Spec{
		Source:     "apache",
		Extensions: []string{"mka"},
	},
	"application/mediaservercontrol+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mscml"},
	},
	"application/vnd.google-apps.document": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"gdoc"},
	},
	"application/vnd.epson.esf": &Spec{
		Source:     "iana",
		Extensions: []string{"esf"},
	},
	"application/x-dtbncx+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"ncx"},
	},
	"application/vnd.fuzzysheet": &Spec{
		Source:     "iana",
		Extensions: []string{"fzs"},
	},
	"image/g3fax": &Spec{
		Source:     "iana",
		Extensions: []string{"g3"},
	},
	"image/vnd.adobe.photoshop": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"psd"},
	},
	"application/vnd.sun.xml.math": &Spec{
		Source:     "apache",
		Extensions: []string{"sxm"},
	},
	"application/x-texinfo": &Spec{
		Source:     "apache",
		Extensions: []string{"texinfo", "texi"},
	},
	"image/vnd.ms-photo": &Spec{
		Source:     "apache",
		Extensions: []string{"wdp"},
	},
	"application/vnd.intu.qfx": &Spec{
		Source:     "iana",
		Extensions: []string{"qfx"},
	},
	"application/x-virtualbox-ova": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"ova"},
	},
	"model/iges": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"igs", "iges"},
	},
	"application/xop+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xop"},
	},
	"model/vnd.opengex": &Spec{
		Source:     "iana",
		Extensions: []string{"ogex"},
	},
	"application/xenc+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xenc"},
	},
	"application/vnd.groove-injector": &Spec{
		Source:     "iana",
		Extensions: []string{"grv"},
	},
	"audio/mp4": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"m4a", "mp4a"},
	},
	"image/vnd.microsoft.icon": &Spec{
		Source:     "iana",
		Extensions: []string{"ico"},
	},
	"application/vnd.solent.sdkm+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"sdkm", "sdkd"},
	},
	"image/x-pict": &Spec{
		Source:     "apache",
		Extensions: []string{"pic", "pct"},
	},
	"application/vnd.symbian.install": &Spec{
		Source:     "apache",
		Extensions: []string{"sis", "sisx"},
	},
	"application/vnd.handheld-entertainment+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"zmm"},
	},
	"application/vnd.route66.link66+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"link66"},
	},
	"application/x-httpd-php": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"php"},
	},
	"application/vnd.ms-word.document.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"docm"},
	},
	"application/vnd.yamaha.hv-script": &Spec{
		Source:     "iana",
		Extensions: []string{"hvs"},
	},
	"video/x-ms-vob": &Spec{
		Source:     "apache",
		Extensions: []string{"vob"},
	},
	"application/vnd.antix.game-component": &Spec{
		Source:     "iana",
		Extensions: []string{"atx"},
	},
	"application/hyperstudio": &Spec{
		Source:     "iana",
		Extensions: []string{"stk"},
	},
	"application/sbml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"sbml"},
	},
	"application/vnd.groove-account": &Spec{
		Source:     "iana",
		Extensions: []string{"gac"},
	},
	"image/vnd.fpx": &Spec{
		Source:     "iana",
		Extensions: []string{"fpx"},
	},
	"application/vnd.anser-web-certificate-issue-initiation": &Spec{
		Source:     "iana",
		Extensions: []string{"cii"},
	},
	"application/vnd.mediastation.cdkey": &Spec{
		Source:     "iana",
		Extensions: []string{"cdkey"},
	},
	"application/x-cdlink": &Spec{
		Source:     "apache",
		Extensions: []string{"vcd"},
	},
	"application/oxps": &Spec{
		Source:     "iana",
		Extensions: []string{"oxps"},
	},
	"message/disposition-notification": &Spec{
		Source:     "iana",
		Extensions: []string{"disposition-notification"},
	},
	"application/vnd.mobius.mbk": &Spec{
		Source:     "iana",
		Extensions: []string{"mbk"},
	},
	"audio/webm": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"weba"},
	},
	"image/x-tga": &Spec{
		Source:     "apache",
		Extensions: []string{"tga"},
	},
	"application/vnd.google-apps.spreadsheet": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"gsheet"},
	},
	"application/pkcs7-signature": &Spec{
		Source:     "iana",
		Extensions: []string{"p7s"},
	},
	"image/x-3ds": &Spec{
		Source:     "apache",
		Extensions: []string{"3ds"},
	},
	"application/rls-services+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rs"},
	},
	"application/vnd.oasis.opendocument.text-web": &Spec{
		Source:     "iana",
		Extensions: []string{"oth"},
	},
	"application/vnd.syncml+xml": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"xsm"},
	},
	"audio/vnd.nuera.ecelp4800": &Spec{
		Source:     "iana",
		Extensions: []string{"ecelp4800"},
	},
	"application/x-dvi": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"dvi"},
	},
	"application/vnd.contact.cmsg": &Spec{
		Source:     "iana",
		Extensions: []string{"cdbcmsg"},
	},
	"image/vnd.tencent.tap": &Spec{
		Source:     "iana",
		Extensions: []string{"tap"},
	},
	"application/vnd.ibm.secure-container": &Spec{
		Source:     "iana",
		Extensions: []string{"sc"},
	},
	"chemical/x-cdx": &Spec{
		Source:     "apache",
		Extensions: []string{"cdx"},
	},
	"application/x-ace-compressed": &Spec{
		Source:     "apache",
		Extensions: []string{"ace"},
	},
	"video/x-ms-wvx": &Spec{
		Source:     "apache",
		Extensions: []string{"wvx"},
	},
	"image/hsj2": &Spec{
		Source:     "iana",
		Extensions: []string{"hsj2"},
	},
	"image/x-pcx": &Spec{
		Source:     "apache",
		Extensions: []string{"pcx"},
	},
	"model/3mf": &Spec{
		Source:     "iana",
		Extensions: []string{"3mf"},
	},
	"application/vnd.kde.kpresenter": &Spec{
		Source:     "iana",
		Extensions: []string{"kpr", "kpt"},
	},
	"image/vnd.valve.source.texture": &Spec{
		Source:     "iana",
		Extensions: []string{"vtf"},
	},
	"application/x-bzip": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"bz"},
	},
	"application/x-chrome-extension": &Spec{
		Extensions: []string{"crx"},
	},
	"application/jsonml+json": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"jsonml"},
	},
	"application/msword": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"doc", "dot"},
	},
	"application/vnd.businessobjects": &Spec{
		Source:     "iana",
		Extensions: []string{"rep"},
	},
	"application/x-msterminal": &Spec{
		Source:     "apache",
		Extensions: []string{"trm"},
	},
	"application/vnd.dece.unspecified": &Spec{
		Source:     "iana",
		Extensions: []string{"uvx", "uvvx"},
	},
	"application/ttml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ttml"},
	},
	"application/vnd.epson.quickanime": &Spec{
		Source:     "iana",
		Extensions: []string{"qam"},
	},
	"text/vnd.in3d.3dml": &Spec{
		Source:     "iana",
		Extensions: []string{"3dml"},
	},
	"application/davmount+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"davmount"},
	},
	"application/pdf": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"pdf"},
	},
	"application/vnd.dolby.mlp": &Spec{
		Source:     "apache",
		Extensions: []string{"mlp"},
	},
	"application/vnd.isac.fcs": &Spec{
		Source:     "iana",
		Extensions: []string{"fcs"},
	},
	"application/x-gramps-xml": &Spec{
		Source:     "apache",
		Extensions: []string{"gramps"},
	},
	"application/xml-dtd": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"dtd"},
	},
	"model/mtl": &Spec{
		Source:     "iana",
		Extensions: []string{"mtl"},
	},
	"application/vnd.ms-outlook": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"msg"},
	},
	"application/vnd.mozilla.xul+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xul"},
	},
	"application/x-blorb": &Spec{
		Source:     "apache",
		Extensions: []string{"blb", "blorb"},
	},
	"application/scvp-cv-response": &Spec{
		Source:     "iana",
		Extensions: []string{"scs"},
	},
	"application/sparql-query": &Spec{
		Source:     "iana",
		Extensions: []string{"rq"},
	},
	"application/vnd.ds-keypoint": &Spec{
		Source:     "apache",
		Extensions: []string{"kpxx"},
	},
	"application/vnd.ms-officetheme": &Spec{
		Source:     "iana",
		Extensions: []string{"thmx"},
	},
	"application/vnd.sailingtracker.track": &Spec{
		Source:     "iana",
		Extensions: []string{"st"},
	},
	"application/x-tex-tfm": &Spec{
		Source:     "apache",
		Extensions: []string{"tfm"},
	},
	"application/vnd.visionary": &Spec{
		Source:     "iana",
		Extensions: []string{"vis"},
	},
	"application/vnd.ms-excel": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"xls", "xlm", "xla", "xlc", "xlt", "xlw"},
	},
	"application/vnd.nokia.n-gage.ac+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ac"},
	},
	"model/vnd.vtu": &Spec{
		Source:     "iana",
		Extensions: []string{"vtu"},
	},
	"image/vnd.djvu": &Spec{
		Source:     "iana",
		Extensions: []string{"djvu", "djv"},
	},
	"application/vnd.tcpdump.pcap": &Spec{
		Source:     "iana",
		Extensions: []string{"pcap", "cap", "dmp"},
	},
	"application/mp21": &Spec{
		Source:     "iana",
		Extensions: []string{"m21", "mp21"},
	},
	"application/vnd.spotfire.sfs": &Spec{
		Source:     "iana",
		Extensions: []string{"sfs"},
	},
	"application/x-arj": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"arj"},
	},
	"chemical/x-cml": &Spec{
		Source:     "apache",
		Extensions: []string{"cml"},
	},
	"text/cache-manifest": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"appcache", "manifest"},
	},
	"application/vnd.anser-web-funds-transfer-initiation": &Spec{
		Source:     "apache",
		Extensions: []string{"fti"},
	},
	"image/apng": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"apng"},
	},
	"application/vnd.lotus-approach": &Spec{
		Source:     "iana",
		Extensions: []string{"apr"},
	},
	"application/vnd.yamaha.smaf-audio": &Spec{
		Source:     "iana",
		Extensions: []string{"saf"},
	},
	"application/x-ms-xbap": &Spec{
		Source:     "apache",
		Extensions: []string{"xbap"},
	},
	"x-conference/x-cooltalk": &Spec{
		Source:     "apache",
		Extensions: []string{"ice"},
	},
	"application/x-glulx": &Spec{
		Source:     "apache",
		Extensions: []string{"ulx"},
	},
	"video/jpeg": &Spec{
		Source:     "iana",
		Extensions: []string{"jpgv"},
	},
	"image/png": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"png"},
	},
	"application/vnd.fdsn.seed": &Spec{
		Source:     "iana",
		Extensions: []string{"seed", "dataless"},
	},
	"application/vnd.triscape.mxs": &Spec{
		Source:     "iana",
		Extensions: []string{"mxs"},
	},
	"application/x-java-archive-diff": &Spec{
		Source:     "nginx",
		Extensions: []string{"jardiff"},
	},
	"application/scvp-vp-request": &Spec{
		Source:     "iana",
		Extensions: []string{"spq"},
	},
	"model/gltf-binary": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"glb"},
	},
	"application/timestamped-data": &Spec{
		Source:     "iana",
		Extensions: []string{"tsd"},
	},
	"application/x-msdos-program": &Spec{
		Extensions: []string{"exe"},
	},
	"application/vnd.syncml.dmddf+xml": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"ddf"},
	},
	"application/vnd.ms-lrm": &Spec{
		Source:     "iana",
		Extensions: []string{"lrm"},
	},
	"image/vnd.airzip.accelerator.azv": &Spec{
		Source:     "iana",
		Extensions: []string{"azv"},
	},
	"application/vnd.wordperfect": &Spec{
		Source:     "iana",
		Extensions: []string{"wpd"},
	},
	"application/voicexml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"vxml"},
	},
	"image/sgi": &Spec{
		Source:     "apache",
		Extensions: []string{"sgi"},
	},
	"application/atsc-held+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"held"},
	},
	"image/jxs": &Spec{
		Source:     "iana",
		Extensions: []string{"jxs"},
	},
	"message/global": &Spec{
		Source:     "iana",
		Extensions: []string{"u8msg"},
	},
	"video/3gpp": &Spec{
		Source:     "iana",
		Extensions: []string{"3gp", "3gpp"},
	},
	"application/vnd.fsc.weblaunch": &Spec{
		Source:     "iana",
		Extensions: []string{"fsc"},
	},
	"application/vnd.nitf": &Spec{
		Source:     "iana",
		Extensions: []string{"ntf", "nitf"},
	},
	"application/vnd.mobius.txf": &Spec{
		Source:     "iana",
		Extensions: []string{"txf"},
	},
	"application/cu-seeme": &Spec{
		Source:     "apache",
		Extensions: []string{"cu"},
	},
	"image/heif-sequence": &Spec{
		Source:     "iana",
		Extensions: []string{"heifs"},
	},
	"application/vnd.tao.intent-module-archive": &Spec{
		Source:     "iana",
		Extensions: []string{"tao"},
	},
	"model/stl": &Spec{
		Source:     "iana",
		Extensions: []string{"stl"},
	},
	"model/vnd.parasolid.transmit.binary": &Spec{
		Source:     "iana",
		Extensions: []string{"x_b"},
	},
	"video/vnd.dece.mobile": &Spec{
		Source:     "iana",
		Extensions: []string{"uvm", "uvvm"},
	},
	"video/x-fli": &Spec{
		Source:     "apache",
		Extensions: []string{"fli"},
	},
	"application/vnd.jcp.javame.midlet-rms": &Spec{
		Source:     "iana",
		Extensions: []string{"rms"},
	},
	"audio/basic": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"au", "snd"},
	},
	"application/x-java-jnlp-file": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"jnlp"},
	},
	"application/vnd.ms-excel.sheet.binary.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"xlsb"},
	},
	"video/x-m4v": &Spec{
		Source:     "apache",
		Extensions: []string{"m4v"},
	},
	"application/java-vm": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"class"},
	},
	"application/x-authorware-map": &Spec{
		Source:     "apache",
		Extensions: []string{"aam"},
	},
	"text/x-asm": &Spec{
		Source:     "apache",
		Extensions: []string{"s", "asm"},
	},
	"application/vnd.hp-hps": &Spec{
		Source:     "iana",
		Extensions: []string{"hps"},
	},
	"application/x-keepass2": &Spec{
		Extensions: []string{"kdbx"},
	},
	"application/ssml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ssml"},
	},
	"application/vnd.hhe.lesson-player": &Spec{
		Source:     "iana",
		Extensions: []string{"les"},
	},
	"application/vnd.citationstyles.style+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"csl"},
	},
	"application/yin+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"yin"},
	},
	"application/vnd.dreamfactory": &Spec{
		Source:     "iana",
		Extensions: []string{"dfac"},
	},
	"application/xslt+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xslt"},
	},
	"chemical/x-cif": &Spec{
		Source:     "apache",
		Extensions: []string{"cif"},
	},
	"image/vnd.xiff": &Spec{
		Source:     "iana",
		Extensions: []string{"xif"},
	},
	"application/vnd.igloader": &Spec{
		Source:     "iana",
		Extensions: []string{"igl"},
	},
	"model/obj": &Spec{
		Source:     "iana",
		Extensions: []string{"obj"},
	},
	"application/x-font-pcf": &Spec{
		Source:     "apache",
		Extensions: []string{"pcf"},
	},
	"application/java-serialized-object": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"ser"},
	},
	"application/pls+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"pls"},
	},
	"image/fits": &Spec{
		Source:     "iana",
		Extensions: []string{"fits"},
	},
	"application/vnd.3gpp.pic-bw-small": &Spec{
		Source:     "iana",
		Extensions: []string{"psb"},
	},
	"application/vnd.fdsn.mseed": &Spec{
		Source:     "iana",
		Extensions: []string{"mseed"},
	},
	"application/vnd.hp-jlyt": &Spec{
		Source:     "iana",
		Extensions: []string{"jlt"},
	},
	"application/vnd.ms-pki.seccat": &Spec{
		Source:     "apache",
		Extensions: []string{"cat"},
	},
	"image/x-mrsid-image": &Spec{
		Source:     "apache",
		Extensions: []string{"sid"},
	},
	"image/x-portable-graymap": &Spec{
		Source:     "apache",
		Extensions: []string{"pgm"},
	},
	"application/n-triples": &Spec{
		Source:     "iana",
		Extensions: []string{"nt"},
	},
	"application/p2p-overlay+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"relo"},
	},
	"application/thraud+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"tfi"},
	},
	"application/vnd.oasis.opendocument.chart": &Spec{
		Source:     "iana",
		Extensions: []string{"odc"},
	},
	"video/mpeg": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"mpeg", "mpg", "mpe", "m1v", "m2v"},
	},
	"image/x-cmx": &Spec{
		Source:     "apache",
		Extensions: []string{"cmx"},
	},
	"application/metalink+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"metalink"},
	},
	"video/x-sgi-movie": &Spec{
		Source:     "apache",
		Extensions: []string{"movie"},
	},
	"application/pkix-cert": &Spec{
		Source:     "iana",
		Extensions: []string{"cer"},
	},
	"text/x-pascal": &Spec{
		Source:     "apache",
		Extensions: []string{"p", "pas"},
	},
	"video/vnd.dece.hd": &Spec{
		Source:     "iana",
		Extensions: []string{"uvh", "uvvh"},
	},
	"video/x-smv": &Spec{
		Source:     "apache",
		Extensions: []string{"smv"},
	},
	"application/vnd.kde.karbon": &Spec{
		Source:     "iana",
		Extensions: []string{"karbon"},
	},
	"application/pkcs7-mime": &Spec{
		Source:     "iana",
		Extensions: []string{"p7m", "p7c"},
	},
	"application/vnd.sun.xml.draw": &Spec{
		Source:     "apache",
		Extensions: []string{"sxd"},
	},
	"application/x-zmachine": &Spec{
		Source:     "apache",
		Extensions: []string{"z1", "z2", "z3", "z4", "z5", "z6", "z7", "z8"},
	},
	"application/vnd.grafeq": &Spec{
		Source:     "iana",
		Extensions: []string{"gqf", "gqs"},
	},
	"application/vnd.mobius.msl": &Spec{
		Source:     "iana",
		Extensions: []string{"msl"},
	},
	"application/postscript": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"ai", "eps", "ps"},
	},
	"application/vnd.kde.kformula": &Spec{
		Source:     "iana",
		Extensions: []string{"kfo"},
	},
	"application/pkcs10": &Spec{
		Source:     "iana",
		Extensions: []string{"p10"},
	},
	"application/x-bdoc": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"bdoc"},
	},
	"text/vnd.in3d.spot": &Spec{
		Source:     "iana",
		Extensions: []string{"spot"},
	},
	"application/x-ms-wmz": &Spec{
		Source:     "apache",
		Extensions: []string{"wmz"},
	},
	"application/lgr+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"lgr"},
	},
	"application/vnd.intercon.formnet": &Spec{
		Source:     "iana",
		Extensions: []string{"xpw", "xpx"},
	},
	"application/x-ustar": &Spec{
		Source:     "apache",
		Extensions: []string{"ustar"},
	},
	"audio/x-ms-wma": &Spec{
		Source:     "apache",
		Extensions: []string{"wma"},
	},
	"application/cdmi-domain": &Spec{
		Source:     "iana",
		Extensions: []string{"cdmid"},
	},
	"application/vnd.fujixerox.ddd": &Spec{
		Source:     "iana",
		Extensions: []string{"ddd"},
	},
	"application/vnd.sun.wadl+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"wadl"},
	},
	"application/x-gtar": &Spec{
		Source:     "apache",
		Extensions: []string{"gtar"},
	},
	"audio/x-aac": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"aac"},
	},
	"application/x-lzh-compressed": &Spec{
		Source:     "apache",
		Extensions: []string{"lzh", "lha"},
	},
	"application/vnd.ezpix-package": &Spec{
		Source:     "iana",
		Extensions: []string{"ez3"},
	},
	"application/vnd.iccprofile": &Spec{
		Source:     "iana",
		Extensions: []string{"icc", "icm"},
	},
	"audio/wave": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"wav"},
	},
	"image/heif": &Spec{
		Source:     "iana",
		Extensions: []string{"heif"},
	},
	"application/vnd.cups-ppd": &Spec{
		Source:     "iana",
		Extensions: []string{"ppd"},
	},
	"application/vnd.noblenet-directory": &Spec{
		Source:     "iana",
		Extensions: []string{"nnd"},
	},
	"application/vnd.accpac.simply.imp": &Spec{
		Source:     "iana",
		Extensions: []string{"imp"},
	},
	"application/vnd.immervision-ivu": &Spec{
		Source:     "iana",
		Extensions: []string{"ivu"},
	},
	"application/vnd.aristanetworks.swi": &Spec{
		Source:     "iana",
		Extensions: []string{"swi"},
	},
	"application/vnd.ms-project": &Spec{
		Source:     "iana",
		Extensions: []string{"mpp", "mpt"},
	},
	"application/vnd.nokia.radio-preset": &Spec{
		Source:     "iana",
		Extensions: []string{"rpst"},
	},
	"application/vnd.osgeo.mapguide.package": &Spec{
		Source:     "iana",
		Extensions: []string{"mgp"},
	},
	"application/x-gca-compressed": &Spec{
		Source:     "apache",
		Extensions: []string{"gca"},
	},
	"application/sensml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"sensmlx"},
	},
	"application/vnd.zzazz.deck+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"zaz"},
	},
	"application/x-shockwave-flash": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"swf"},
	},
	"application/x-xfig": &Spec{
		Source:     "apache",
		Extensions: []string{"fig"},
	},
	"image/x-ms-bmp": &Spec{
		Source:       "nginx",
		Compressible: &trueVal,
		Extensions:   []string{"bmp"},
	},
	"application/vnd.stepmania.stepchart": &Spec{
		Source:     "iana",
		Extensions: []string{"sm"},
	},
	"application/x-stuffitx": &Spec{
		Source:     "apache",
		Extensions: []string{"sitx"},
	},
	"application/vnd.macports.portpkg": &Spec{
		Source:     "iana",
		Extensions: []string{"portpkg"},
	},
	"application/vnd.stardivision.draw": &Spec{
		Source:     "apache",
		Extensions: []string{"sda"},
	},
	"application/x-abiword": &Spec{
		Source:     "apache",
		Extensions: []string{"abw"},
	},
	"image/x-rgb": &Spec{
		Source:     "apache",
		Extensions: []string{"rgb"},
	},
	"application/mods+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mods"},
	},
	"application/vnd.epson.salt": &Spec{
		Source:     "iana",
		Extensions: []string{"slt"},
	},
	"application/x-tex": &Spec{
		Source:     "apache",
		Extensions: []string{"tex"},
	},
	"application/x-pkcs12": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"p12", "pfx"},
	},
	"image/cgm": &Spec{
		Source:     "iana",
		Extensions: []string{"cgm"},
	},
	"application/xcap-el+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xel"},
	},
	"video/h261": &Spec{
		Source:     "iana",
		Extensions: []string{"h261"},
	},
	"application/vnd.fujixerox.docuworks.binder": &Spec{
		Source:     "iana",
		Extensions: []string{"xbd"},
	},
	"application/vnd.uiq.theme": &Spec{
		Source:     "iana",
		Extensions: []string{"utz"},
	},
	"application/x-authorware-bin": &Spec{
		Source:     "apache",
		Extensions: []string{"aab", "x32", "u32", "vox"},
	},
	"image/vnd.fujixerox.edmics-rlc": &Spec{
		Source:     "iana",
		Extensions: []string{"rlc"},
	},
	"application/mmt-aei+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"maei"},
	},
	"application/vnd.picsel": &Spec{
		Source:     "iana",
		Extensions: []string{"efif"},
	},
	"application/vnd.software602.filler.form+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"fo"},
	},
	"application/vnd.wt.stf": &Spec{
		Source:     "iana",
		Extensions: []string{"stf"},
	},
	"application/dssc+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xdssc"},
	},
	"application/dash+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"mpd"},
	},
	"application/vnd.mfer": &Spec{
		Source:     "iana",
		Extensions: []string{"mwf"},
	},
	"image/vnd.fst": &Spec{
		Source:     "iana",
		Extensions: []string{"fst"},
	},
	"application/vnd.mseq": &Spec{
		Source:     "iana",
		Extensions: []string{"mseq"},
	},
	"text/vnd.wap.wmlscript": &Spec{
		Source:     "iana",
		Extensions: []string{"wmls"},
	},
	"text/vtt": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Compressible: &trueVal,
		Extensions:   []string{"vtt"},
	},
	"application/x-lua-bytecode": &Spec{
		Extensions: []string{"luac"},
	},
	"image/wmf": &Spec{
		Source:     "iana",
		Extensions: []string{"wmf"},
	},
	"text/vnd.curl.scurl": &Spec{
		Source:     "apache",
		Extensions: []string{"scurl"},
	},
	"application/marc": &Spec{
		Source:     "iana",
		Extensions: []string{"mrc"},
	},
	"application/vnd.hp-pclxl": &Spec{
		Source:     "iana",
		Extensions: []string{"pclxl"},
	},
	"video/webm": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"webm"},
	},
	"application/vnd.ms-htmlhelp": &Spec{
		Source:     "iana",
		Extensions: []string{"chm"},
	},
	"application/vnd.stardivision.impress": &Spec{
		Source:     "apache",
		Extensions: []string{"sdd"},
	},
	"application/vnd.google-earth.kml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"kml"},
	},
	"application/vnd.ms-powerpoint.slideshow.macroenabled.12": &Spec{
		Source:     "iana",
		Extensions: []string{"ppsm"},
	},
	"application/xspf+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"xspf"},
	},
	"application/vnd.semd": &Spec{
		Source:     "iana",
		Extensions: []string{"semd"},
	},
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"docx"},
	},
	"video/quicktime": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"qt", "mov"},
	},
	"application/wasm": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"wasm"},
	},
	"application/x-msdownload": &Spec{
		Source:     "apache",
		Extensions: []string{"exe", "dll", "com", "bat", "msi"},
	},
	"audio/vnd.nuera.ecelp7470": &Spec{
		Source:     "iana",
		Extensions: []string{"ecelp7470"},
	},
	"application/vnd.geospace": &Spec{
		Source:     "iana",
		Extensions: []string{"g3w"},
	},
	"audio/x-mpegurl": &Spec{
		Source:     "apache",
		Extensions: []string{"m3u"},
	},
	"application/x-tar": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"tar"},
	},
	"message/global-headers": &Spec{
		Source:     "iana",
		Extensions: []string{"u8hdr"},
	},
	"audio/vnd.lucent.voice": &Spec{
		Source:     "iana",
		Extensions: []string{"lvp"},
	},
	"application/xaml+xml": &Spec{
		Source:       "apache",
		Compressible: &trueVal,
		Extensions:   []string{"xaml"},
	},
	"image/jp2": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"jp2", "jpg2"},
	},
	"text/x-vcard": &Spec{
		Source:     "apache",
		Extensions: []string{"vcf"},
	},
	"video/vnd.ms-playready.media.pyv": &Spec{
		Source:     "iana",
		Extensions: []string{"pyv"},
	},
	"application/sparql-results+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"srx"},
	},
	"application/pgp-signature": &Spec{
		Source:     "iana",
		Extensions: []string{"asc", "sig"},
	},
	"application/vnd.mophun.certificate": &Spec{
		Source:     "iana",
		Extensions: []string{"mpc"},
	},
	"application/vnd.chemdraw+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"cdxml"},
	},
	"application/vnd.ezpix-album": &Spec{
		Source:     "iana",
		Extensions: []string{"ez2"},
	},
	"application/vnd.oasis.opendocument.presentation-template": &Spec{
		Source:     "iana",
		Extensions: []string{"otp"},
	},
	"application/vnd.rn-realmedia": &Spec{
		Source:     "apache",
		Extensions: []string{"rm"},
	},
	"audio/vnd.dra": &Spec{
		Source:     "iana",
		Extensions: []string{"dra"},
	},
	"application/vnd.ibm.modcap": &Spec{
		Source:     "iana",
		Extensions: []string{"afp", "listafp", "list3820"},
	},
	"application/vnd.trid.tpt": &Spec{
		Source:     "iana",
		Extensions: []string{"tpt"},
	},
	"application/x-7z-compressed": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"7z"},
	},
	"image/x-portable-bitmap": &Spec{
		Source:     "apache",
		Extensions: []string{"pbm"},
	},
	"application/vnd.yamaha.openscoreformat": &Spec{
		Source:     "iana",
		Extensions: []string{"osf"},
	},
	"application/json5": &Spec{
		Extensions: []string{"json5"},
	},
	"application/vnd.frogans.ltf": &Spec{
		Source:     "iana",
		Extensions: []string{"ltf"},
	},
	"application/x-authorware-seg": &Spec{
		Source:     "apache",
		Extensions: []string{"aas"},
	},
	"application/x-ms-shortcut": &Spec{
		Source:     "apache",
		Extensions: []string{"lnk"},
	},
	"image/x-jng": &Spec{
		Source:     "nginx",
		Extensions: []string{"jng"},
	},
	"application/senml+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"senmlx"},
	},
	"application/x-director": &Spec{
		Source:     "apache",
		Extensions: []string{"dir", "dcr", "dxr", "cst", "cct", "cxt", "w3d", "fgd", "swa"},
	},
	"application/toml": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"toml"},
	},
	"text/yaml": &Spec{
		Extensions: []string{"yaml", "yml"},
	},
	"application/n-quads": &Spec{
		Source:     "iana",
		Extensions: []string{"nq"},
	},
	"text/turtle": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Extensions:   []string{"ttl"},
	},
	"application/mrb-publish+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xdf"},
	},
	"application/vnd.ms-works": &Spec{
		Source:     "iana",
		Extensions: []string{"wps", "wks", "wcm", "wdb"},
	},
	"font/collection": &Spec{
		Source:     "iana",
		Extensions: []string{"ttc"},
	},
	"video/x-flv": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"flv"},
	},
	"application/x-xpinstall": &Spec{
		Source:       "apache",
		Compressible: &falseVal,
		Extensions:   []string{"xpi"},
	},
	"chemical/x-csml": &Spec{
		Source:     "apache",
		Extensions: []string{"csml"},
	},
	"application/vnd.fujixerox.docuworks": &Spec{
		Source:     "iana",
		Extensions: []string{"xdw"},
	},
	"application/vnd.unity": &Spec{
		Source:     "iana",
		Extensions: []string{"unityweb"},
	},
	"image/jpeg": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"jpeg", "jpg", "jpe"},
	},
	"application/x-chess-pgn": &Spec{
		Source:     "apache",
		Extensions: []string{"pgn"},
	},
	"application/cdmi-object": &Spec{
		Source:     "iana",
		Extensions: []string{"cdmio"},
	},
	"application/xcap-ns+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xns"},
	},
	"application/vnd.sun.xml.calc": &Spec{
		Source:     "apache",
		Extensions: []string{"sxc"},
	},
	"application/x-virtualbox-vmdk": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"vmdk"},
	},
	"audio/vnd.dts.hd": &Spec{
		Source:     "iana",
		Extensions: []string{"dtshd"},
	},
	"image/ief": &Spec{
		Source:     "iana",
		Extensions: []string{"ief"},
	},
	"application/cdmi-capability": &Spec{
		Source:     "iana",
		Extensions: []string{"cdmia"},
	},
	"text/prs.lines.tag": &Spec{
		Source:     "iana",
		Extensions: []string{"dsc"},
	},
	"text/x-sass": &Spec{
		Extensions: []string{"sass"},
	},
	"application/vnd.kde.kivio": &Spec{
		Source:     "iana",
		Extensions: []string{"flw"},
	},
	"application/x-wais-source": &Spec{
		Source:     "apache",
		Extensions: []string{"src"},
	},
	"application/vnd.wap.wmlscriptc": &Spec{
		Source:     "iana",
		Extensions: []string{"wmlsc"},
	},
	"application/x-envoy": &Spec{
		Source:     "apache",
		Extensions: []string{"evy"},
	},
	"chemical/x-cmdf": &Spec{
		Source:     "apache",
		Extensions: []string{"cmdf"},
	},
	"application/reginfo+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"rif"},
	},
	"application/vnd.apple.keynote": &Spec{
		Source:     "iana",
		Extensions: []string{"key"},
	},
	"text/sgml": &Spec{
		Source:     "iana",
		Extensions: []string{"sgml", "sgm"},
	},
	"application/x-debian-package": &Spec{
		Source:     "apache",
		Extensions: []string{"deb", "udeb"},
	},
	"application/pkix-crl": &Spec{
		Source:     "iana",
		Extensions: []string{"crl"},
	},
	"application/pkixcmp": &Spec{
		Source:     "iana",
		Extensions: []string{"pki"},
	},
	"text/vnd.sun.j2me.app-descriptor": &Spec{
		Source:       "iana",
		CharEncoding: "UTF-8",
		Extensions:   []string{"jad"},
	},
	"application/vnd.crick.clicker.palette": &Spec{
		Source:     "iana",
		Extensions: []string{"clkp"},
	},
	"application/cdmi-queue": &Spec{
		Source:     "iana",
		Extensions: []string{"cdmiq"},
	},
	"application/vnd.neurolanguage.nlu": &Spec{
		Source:     "iana",
		Extensions: []string{"nlu"},
	},
	"application/x-research-info-systems": &Spec{
		Source:     "apache",
		Extensions: []string{"ris"},
	},
	"image/x-cmu-raster": &Spec{
		Source:     "apache",
		Extensions: []string{"ras"},
	},
	"text/shex": &Spec{
		Extensions: []string{"shex"},
	},
	"text/x-lua": &Spec{
		Extensions: []string{"lua"},
	},
	"application/vnd.seemail": &Spec{
		Source:     "iana",
		Extensions: []string{"see"},
	},
	"application/cdmi-container": &Spec{
		Source:     "iana",
		Extensions: []string{"cdmic"},
	},
	"application/vnd.pvi.ptid1": &Spec{
		Source:     "iana",
		Extensions: []string{"ptid"},
	},
	"application/vnd.stepmania.package": &Spec{
		Source:     "iana",
		Extensions: []string{"smzip"},
	},
	"application/x-pilot": &Spec{
		Source:     "nginx",
		Extensions: []string{"prc", "pdb"},
	},
	"application/tei+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"tei", "teicorpus"},
	},
	"application/vnd.sema": &Spec{
		Source:     "iana",
		Extensions: []string{"sema"},
	},
	"application/x-sv4cpio": &Spec{
		Source:     "apache",
		Extensions: []string{"sv4cpio"},
	},
	"application/vnd.kodak-descriptor": &Spec{
		Source:     "iana",
		Extensions: []string{"sse"},
	},
	"application/vnd.pocketlearn": &Spec{
		Source:     "iana",
		Extensions: []string{"plf"},
	},
	"application/vnd.wolfram.player": &Spec{
		Source:     "iana",
		Extensions: []string{"nbp"},
	},
	"application/vnd.fujitsu.oasysprs": &Spec{
		Source:     "iana",
		Extensions: []string{"bh2"},
	},
	"application/vnd.oasis.opendocument.database": &Spec{
		Source:     "iana",
		Extensions: []string{"odb"},
	},
	"application/vnd.stardivision.math": &Spec{
		Source:     "apache",
		Extensions: []string{"smf"},
	},
	"audio/mp3": &Spec{
		Compressible: &falseVal,
		Extensions:   []string{"mp3"},
	},
	"audio/ogg": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"oga", "ogg", "spx"},
	},
	"application/x-redhat-package-manager": &Spec{
		Source:     "nginx",
		Extensions: []string{"rpm"},
	},
	"application/vnd.micrografx.igx": &Spec{
		Source:     "iana",
		Extensions: []string{"igx"},
	},
	"image/jls": &Spec{
		Source:     "iana",
		Extensions: []string{"jls"},
	},
	"application/vnd.mobius.plc": &Spec{
		Source:     "iana",
		Extensions: []string{"plc"},
	},
	"application/vnd.mophun.application": &Spec{
		Source:     "iana",
		Extensions: []string{"mpn"},
	},
	"application/vnd.smaf": &Spec{
		Source:     "iana",
		Extensions: []string{"mmf"},
	},
	"text/spdx": &Spec{
		Source:     "iana",
		Extensions: []string{"spdx"},
	},
	"application/its+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"its"},
	},
	"application/rpki-ghostbusters": &Spec{
		Source:     "iana",
		Extensions: []string{"gbr"},
	},
	"application/vnd.ms-artgalry": &Spec{
		Source:     "iana",
		Extensions: []string{"cil"},
	},
	"application/mmt-usd+xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"musd"},
	},
	"message/global-delivery-status": &Spec{
		Source:     "iana",
		Extensions: []string{"u8dsn"},
	},
	"application/vnd.frogans.fnc": &Spec{
		Source:     "iana",
		Extensions: []string{"fnc"},
	},
	"application/vnd.ms-cab-compressed": &Spec{
		Source:     "iana",
		Extensions: []string{"cab"},
	},
	"application/x-chat": &Spec{
		Source:     "apache",
		Extensions: []string{"chat"},
	},
	"application/scvp-vp-response": &Spec{
		Source:     "iana",
		Extensions: []string{"spp"},
	},
	"application/vnd.ms-fontobject": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"eot"},
	},
	"application/vnd.ibm.minipay": &Spec{
		Source:     "iana",
		Extensions: []string{"mpy"},
	},
	"application/vnd.joost.joda-archive": &Spec{
		Source:     "iana",
		Extensions: []string{"joda"},
	},
	"application/xml": &Spec{
		Source:       "iana",
		Compressible: &trueVal,
		Extensions:   []string{"xml", "xsl", "xsd", "rng"},
	},
	"audio/3gpp": &Spec{
		Source:       "iana",
		Compressible: &falseVal,
		Extensions:   []string{"3gpp"},
	},
	"application/vnd.airzip.filesecure.azs": &Spec{
		Source:     "iana",
		Extensions: []string{"azs"},
	},
	"application/vnd.dece.data": &Spec{
		Source:     "iana",
		Extensions: []string{"uvf", "uvvf", "uvd", "uvvd"},
	},
	"application/vnd.dece.zip": &Spec{
		Source:     "iana",
		Extensions: []string{"uvz", "uvvz"},
	},
	"application/x-virtualbox-vdi": &Spec{
		Compressible: &trueVal,
		Extensions:   []string{"vdi"},
	},
}
