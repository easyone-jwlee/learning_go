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

<hr/>

### Error

* `go fmt ./...` 명령어 작성시 오류 발생

```bash
ch1 > make                                                                             0:07
go fmt ./...
pattern ./...: directory prefix . does not contain main module or its selected dependencies
make: *** [fmt] Error 1
```

* 1.16 이후에 릴리즈 된 Go는 기본적으로 모듈의 형태로 빌드를 시도하여 go.mod를 만들지 않으면 go fmt ./... 하는 시점에 오류가 발생한다. 이를 해결하기 위해서는 `go env -w GO111MODULE=auto`를 명령라인에 입력하여 적용한 뒤 다시 시도하면 된다.
 * 참고: <https://go.dev/blog/go116-module-changes>