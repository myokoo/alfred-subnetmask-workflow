# alfred-subnetmask-workflow

Alfred workflow to calculate subnet mask.

#### subnetmask xxx.xxx.xxx.xxxx/xx or subnetmask xxx.xxx.xxx.xxx/xxx.xxx.xxx.xxx 
ex1) subnetmask 192.168.111.33/24  
ex2) subnetmask 192.185.111.33/255.255.255.0  

![subnetmask](https://user-images.githubusercontent.com/1995330/51617935-9a7a1100-1f70-11e9-905c-3907a05564fc.gif)


#### subnetmask xxx.xxx.xxx.xxx 
ex) subnetmask 255.255.255.0  

![broadcast](https://user-images.githubusercontent.com/1995330/51617904-86ceaa80-1f70-11e9-8fc6-d37d795b6c3a.gif)


#### subnetmask /xx  
ex) subnetmask /24  

![prefix](https://user-images.githubusercontent.com/1995330/51617833-643c9180-1f70-11e9-85f2-f4f50b2b8078.gif)



### Setup

#### go version >= 1.11.1

```bash
git clone git@github.com:myokoo/alfred-subnetmask-workflow.git
GO111MODULE=on go build

sh setup.sh
```

`setup.sh` is created link to alfred workflow v3 directory.

### Licensing
This library is released under the [MIT license][license].

The icon is <a href="http://creativecommons.org/licenses/by/4.0/" rel="license"><img style="border-width: 0; font-size: 15px; vertical-align: middle;" src="https://i.creativecommons.org/l/by/4.0/80x15.png" alt="クリエイティブ・コモンズ・ライセンス" /></a> [SAKURA Internet Inc.](https://knowledge.sakura.ad.jp/4724/)

[license]: ./LICENSE
