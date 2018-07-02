# alfred-subnetmask-workflow

Alfred workflow to calculate subnet mask.

#### subnetmask xxx.xxx.xxx.xxxx/xx
![subnetmask](https://user-images.githubusercontent.com/1995330/42171666-d7fa375c-7e54-11e8-8feb-45159e8eb14f.png)


#### subnetmask xxx.xxx.xxx.xxx
![broadcast](https://user-images.githubusercontent.com/1995330/42171661-d44db584-7e54-11e8-8454-30e457e6940b.png)

#### subnetmask /xx
![prefix](https://user-images.githubusercontent.com/1995330/42171671-da4dce60-7e54-11e8-90de-484ad6ea468a.png)



### Setup

#### go version >= 1.10.3

```bash
go get golang.org/x/vgo
vgo build

sh setup.sh
```

`setup.sh` is created link to alfred workflow v3 directory.

### Licensing
This library is released under the [MIT licence][licence].

The icon is <a href="http://creativecommons.org/licenses/by/4.0/" rel="license"><img style="border-width: 0; font-size: 15px; vertical-align: middle;" src="https://i.creativecommons.org/l/by/4.0/80x15.png" alt="クリエイティブ・コモンズ・ライセンス" /></a> [SAKURA Internet Inc.](https://knowledge.sakura.ad.jp/4724/)

[licence]: ./LICENCE
