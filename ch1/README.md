### 1.3.2 Third-party tools with go

* HTTP 서버의 부하 테스트를 위한 `hey`

```bash
ch1 > go install github.com/rakyll/hey@latest 
ch1 > hey https://www.golang.org                      

Summary:
  Total:        3.5178 secs
  Slowest:      0.0000 secs
  Fastest:      0.0000 secs
  Average:       NaN secs
  Requests/sec: 56.8543
  

Response time histogram:


Latency distribution:

Details (average, fastest, slowest):
  DNS+dialup:    NaN secs, 0.0000 secs, 0.0000 secs
  DNS-lookup:    NaN secs, 0.0000 secs, 0.0000 secs
  req write:     NaN secs, 0.0000 secs, 0.0000 secs
  resp wait:     NaN secs, 0.0000 secs, 0.0000 secs
  resp read:     NaN secs, 0.0000 secs, 0.0000 secs

Status code distribution:

Error distribution:
  [200] Get "https://go.dev/": EOF

```

> `go install`을 이용하여 프로그램을 받을 수 있도록 프로그램을 배포할 필요도 없다. 
> binary를 download 받을 수 있게 하면 되니까. 
> 하지만 `go install`으로 설치할 수 있게 만들면 다른 gopher들이 매우 편리하게 배포 받을 수 있다.

### 1.3.3 Code formatting

* Go의 주요 설계 목표 중 하나는 코드 작성이 효율적인 언어를 만드는 것.
    * 구문이 단순하고 컴파일러가 빨라야 한다는 의미.
* Go 개발 도구에는 표준 포맷으로 코드를 자동으로 맞춰주는 `go fmt` command가 포함되어 있다.
* import문을 정리해주는 `goimports`라는 향상된 버전의 `go fmt`가 있다.
    * import line들을 알파벳순으로 정렬하고 사용되지 않는 import package는 제거한다.
    * 지정되지 않은 import package를 추측해서 추가해주지만, 때로는 정확하지 않아서 직접 import line을 추가해주어야 한다.

```bash
ch1 > make goimports_inst  
go install golang.org/x/tools/cmd/goimports@latest
# -l option은 포맷 수정이 필요한 파일들을 콘솔에 출력.
# -w option은 goimports에게 파일에 직접 수정을 하도록 지시
ch1 > goimports -l -w .      
ch1 >     
```

### 1.4 linting and vetting

* 코드가 관용적이고 좋은 품질을 갖도록 코드의 포맷이 잘 지켜졌는지 확인해야한다.
* [Effective Go][effecttive go link]와 Go의 위키에 있는 [Code Review Comments][Code Review Comments]를 읽고 관용적인 Go 코드가 어떻게 보이는가 이해해야 한다.

[effecttive go link]: https://go.dev/doc/effective_go
[Code Review Comments]: https://github.com/golang/go/wiki/CodeReviewComments

#### go lint

* Code가 style guide를 잘 지켰는지 확인한다. 
* 100% 신뢰할 수는 없기 때문에 golint가 제안하는 것을 진지하게 생각은 하되, 무조건 수용할 필요는 없다.

```bash
ch1 > make golint_inst                                
go install golang.org/x/lint/golint@latest
ch1 > golint ./...       
hello.go:7:6: exported type Param should have comment or be unexported
param_pac/param.go:1:1: don't use an underscore in package name
param_pac/param.go:3:6: exported type Param should have comment or be unexported
param_pac/param.go:7:1: exported method Param.NewParam should have comment or be unexported
```

#### go vet

* Go에 포함된 도구

```bash
ch1 > go vet ./...  
```

#### golangci-lint

* 추가적인 Code Style과 잠재적인 bug를 찾아주는 third-party tool들이 있다.
* golangci-lint tool를 이용해 통합해서 실행할 수 있음.
* [golangci-lint install][golangci-lint install]

[golangci-lint install]: https://golangci-lint.run/usage/install/

```bash
# macOS 를 사용하고 있기때문에 brew를 이용하여 설치
ch1 > brew install golangci-lint   
Running `brew update --auto-update`...
==> Homebrew collects anonymous analytics.
Read the analytics documentation (and how to opt-out) here:
  https://docs.brew.sh/Analytics
No analytics have been recorded yet (nor will be during this `brew` run).

==> Homebrew is run entirely by unpaid volunteers. Please consider donating:
  https://github.com/Homebrew/brew#donations

==> Auto-updated Homebrew!
Updated 2 taps (homebrew/core and homebrew/cask).
==> New Formulae
bacon                        couchbase-shell              fw                           llm                          pixi                         quictls                      strip-nondeterminism         wpscan
bfs                          crystalline                  go-feature-flag              mailpit                      plog                         ruff-lsp                     terragrunt-atlantis-config   yyjson
bilix                        dysk                         govulncheck                  mvfst                        pop                          runme                        tzdiff
blink                        erg                          hoverfly                     mvt                          prettierd                    sh4d0wup                     webpod
botan@2                      espflash                     ittapi                       neonctl                      pylyzer                      solr@8.11                    wget2
cbonsai                      fuc                          killport                     pgrok                        python-cryptography          sqlpage                      woof-doom
==> New Casks
4k-video-downloaderplus   elektron-overbridge       fedistar                  herd                      lm-studio                 mycard                    replay                    screen-studio             showmeyourhotkeys
chatall                   ente                      glaze                     keyclu                    maa                       poe                       ripx                      sfm

==> Fetching dependencies for golangci-lint: go
==> Fetching go
==> Downloading https://ghcr.io/v2/homebrew/core/go/manifests/1.20.6
#################################################################################################################################################################################################################################### 100.0%
==> Downloading https://ghcr.io/v2/homebrew/core/go/blobs/sha256:2bcdf995616b5dce70e7af6ea081de4bcb9c128ee352136ed585ccd6f19201e5
#################################################################################################################################################################################################################################### 100.0%
==> Fetching golangci-lint
==> Downloading https://ghcr.io/v2/homebrew/core/golangci-lint/manifests/1.53.3
#################################################################################################################################################################################################################################### 100.0%
==> Downloading https://ghcr.io/v2/homebrew/core/golangci-lint/blobs/sha256:1342963702b51e520e2ac0b020ca65cd071d00bfc5f9cd62c05fd4182cce2fd6
#################################################################################################################################################################################################################################### 100.0%
==> Installing dependencies for golangci-lint: go
==> Installing golangci-lint dependency: go
==> Pouring go--1.20.6.arm64_ventura.bottle.tar.gz
🍺  /opt/homebrew/Cellar/go/1.20.6: 11,997 files, 233.1MB
==> Installing golangci-lint
==> Pouring golangci-lint--1.53.3.arm64_ventura.bottle.tar.gz
==> Caveats
zsh completions have been installed to:
  /opt/homebrew/share/zsh/site-functions
==> Summary
🍺  /opt/homebrew/Cellar/golangci-lint/1.53.3: 9 files, 29.9MB
==> Running `brew cleanup golangci-lint`...
Disable this behaviour by setting HOMEBREW_NO_INSTALL_CLEANUP.
Hide these hints with HOMEBREW_NO_ENV_HINTS (see `man brew`).
==> Caveats
==> golangci-lint
zsh completions have been installed to:
  /opt/homebrew/share/zsh/site-functions
ch1 > brew upgrade golangci-lint   
Warning: golangci-lint 1.53.3 already installed
ch1 > golangci-lint run 
param_pac/param.go:4:2: field `x` is unused (unused)
        x int
        ^
hello.go:8:2: field `x` is unused (unused)
        x int
        ^
param_pac/param.go:8:2: S1023: redundant `return` statement (gosimple)
        return
        ^
```

* Project의 root directory에 `.golangci.yml` file을 만들어 활성화하고자하는 linter와 분석할 file을 설정할 수 있다.

### 1.6 Makefiles

* `.DEFAULT_GOAL`은 어떤 Target도 지정하지 않았을 때, 기본적으로 수행하는 Target을 정의한다.
* `.PHONY` line은 target과 동일한 이름으로 dir를 생성한 경우, make가 혼동되지 않도록 한다.


<hr/>

### Error

* `go fmt ./...` 명령어 작성시 오류 발생

```bash
ch1 > make                                               
go fmt ./...
pattern ./...: directory prefix . does not contain main module or its selected dependencies
make: *** [fmt] Error 1
```

* 1.16 이후에 릴리즈 된 Go는 기본적으로 모듈의 형태로 빌드를 시도하여 go.mod를 만들지 않으면 go fmt ./... 하는 시점에 오류가 발생한다. 이를 해결하기 위해서는 `go env -w GO111MODULE=auto`를 명령라인에 입력하여 적용한 뒤 다시 시도하면 된다.
 * 참고: <https://go.dev/blog/go116-module-changes>