
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

中文简繁互转  
Chinese-Character Jian<=>Fan converter


# API

#### > {{cat "jianfan_test.go" | color "go"}}

All patches welcome.

# Credits

- [github.com/siongui/gojianfan](https://github.com/siongui/gojianfan)
  where the code was initially ported from
- [OpenCC](https://github.com/BYVoid/OpenCC)
  where the Jian<=>Fan lookup table is based on
