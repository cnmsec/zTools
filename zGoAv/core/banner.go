package core

import (
	"github.com/gookit/color"
	"github.com/zGoAv/gologger"
)

const Version = "0.1"

const Banner = `
                            _  _  _                                      _                               
                         _ (_)(_)(_) _                                 _(_)_                             
      _  _  _  _        (_)         (_)          _  _  _             _(_) (_)_        _               _  
     (_)(_)(_)(_)       (_)    _  _  _        _ (_)(_)(_) _        _(_)     (_)_     (_)_           _(_) 
           _ (_)        (_)   (_)(_)(_)      (_)         (_)      (_) _  _  _ (_)      (_)_       _(_)   
        _ (_)           (_)         (_)      (_)         (_)      (_)(_)(_)(_)(_)        (_)_   _(_)     
      _(_)  _  _        (_) _  _  _ (_)      (_) _  _  _ (_)      (_)         (_)          (_)_(_)       
     (_)(_)(_)(_)          (_)(_)(_)(_)         (_)(_)(_)         (_)         (_)            (_)

"这是一款免杀生成工具，需要用shellcode生成"
""
`
const LinkAndAuthor = "仅用于测试"
const Warning = "警告：\\n1.本工具仅用于企业内部测试，请勿用于任何非法犯罪活动，否则后果自负\\n2.本工具需要Go语言环境，且使用时需要关闭杀软"

func ShowBanner() {
	color.RGBStyleFromString("210,105,30").Println(Banner)
	color.RGBStyleFromString("255,0,0").Println(Warning)
	color.RGBStyleFromString("30,144,255").Println(LinkAndAuthor)
	gologger.Infof("Current Version:%s\n", Version)
}
