Gontainer Toy Project 에 대한 대략적인 기획. 틈틈히 해보자

## 1. Overview
---
  Gontainer는 Container Runtime 이 설치된 환경에 관리가 필요한 Container 의 상태를 UI 로 보여주는 프로젝트이다.
Go 기반의 Backend 와 React 기반의 Frontend 로 구성된다.

- Go 1.23.2
    - web 서버는 별도 Framework(ex. Gin, Echo, Fiber 등) 사용 하지 않고, net/http 를 이용하여 생성
    - API Specification 제공을 위한 Swagger 추가
    - Logger 는 uber 에서 만든 zap 사용. 기존에 많이 사용하던 logurs 대비 성능이 뛰어나고 구조화 로깅(json 형태) 을 지원, File Logging 없이 STDOUT 사용
    - 상태 관리는 최소화. => Persistent Layer 없음. ? 필요 시 SQLite w/ GORM 을 통해 구성.
    - 인증/인가 기능은 필요 시 추후 적용 예정
    - 다른 Node 의 정보를 얻는 방법 설계. 
        - ex. F/E -> B/E -> 다른 Node Call(실시간)
        - F/E -> B/E <- 주기적인 상태 보고
        - F/E -> B/E -/-> 다른 Node 수집(Scrape)
    - Container Log Streaming(Tailing) API?
- React + Vite + TypeScript
    - Client Side Rendering 을 이용함. 가볍고 빠르게 만들기 위함. 정적 빌드를 통해 /static 하위로 넣을 예정임
    - Server Side Rendering 및 무거운 Next,js 와 같은 Framework 은 사용하지 않음. 추후 고려
    - Container 항목을 Tile 형태로 제공.

본 프로젝트 사용 대상 인원은 다음과 같다.
- Docker 명령어에 익숙하지 않은 사람들
- Docker 가 없는, ex. Containerd 만 있는 환경에 익숙하지 않은 사람들
- Docker Desktop 같이 UI 기반으로 Container를 시작/종료만 하고 싶은 사람들

기능 및 화면 구성 참고하자
- Open Lens
- Docker Desktop
- Rancher

## Project Structure
---
Go 기반의 HTTP 서버를 포함한 Backend 와 React + Vite 기반의 Frontend 를 모노리포(Mono Repository) 형태로 아래와 같은 directory 구조로 사용

- api
    - handler
        - ContainerHandler
        - NodeHandler
        - …
    - router
- cmd
    - gontainer
- conf
- docs
- internal
    - service
        - ContainerService
        - NodeService
        - …
    - domain
        - ContainerRepository
        - NodeRepository
        - …
    - logger
- package
    - gontainer-client
- scripts
    - Containerize
    - Build
    - compose.yaml
- web
    - static
    - src
    - package.json


- 구현
    - API
    - UI
    - Swagger


## API Specifications
---

- JSON Type
- Root URL 은 환경 변수로 default. gontainer
### Healthcheck

1. Healthcheck: Application 의 정상 동작 여부 확인., Container Runtime 의 정상 여부를 점검하고 200을 리턴
    1. GET /api/v1/healthy

### 정적 파일 serving
- Go 의 net/http package 에서 제공하는 FileServer 기능을 활용하여 정적 파일(html/js/css 등) 을 서빙
- /web/static 경로를 /gontainer/index.,html 로 서빙

### Container 관리

UNIX Socket 혹은 HTTP/TCP 를 통해 Container 의 목록을 조회하고 관리하는 기능을 가진다
- Containerd: /run/containerd/containerd.sock

1. Container 목록 조회
    1. GET /gontainer/api/v1/containers
    2. GET /gontainer/api/v1/node/:nodeId/containers
2. Container 상세 조회
    1. GET /gontainer/api/v1/container/:containerId
    2. GET /gontainer/api/v1/node/:nodeId/container/:containerId
3. Image 목록 조회
    1. GET /gontainer/api/v1/images
    2. GET /gontainer/api/v1/node/:nodeId/images
4. Image 상세 조회
    1. GET /gontainer/api/v1/image/:imageId
    2. GET /gontainer/api/v1/node/:nodeId/image/:imageId
5. Container 생성
    1. POST /gontainer/api/v1/container
    2. POST /gontainer/api/v1/node/:nodeId/container
    3. Request Body
             {
           “containerName”: “test_container”,
           ”image”: “nginx:1.23.3”,
           “port”: “8080:80”,
       }
    - volume 도 필요한가?
6. Container Start/Stop
7. Container 삭제
    1. DELETE /gontainer/api/v1/container/:containerId
    2. DELETE /gontainer/api/v1/node/:nodeId/container/:containerId
8. Container Commit
9. Container Push


### Node 관리

1. Node Register: Application 기동 시 Host의 정보를 서버로 전달
    1. Hostname: Node ID로 간주
    2. containerd의 info api 활용: OS, Version 등
2. Node 목록 조회
    1. GET /gontainer/api/v1/nodes
3. Node 상세 조회
    1. GET /gontainer/api/v1/node/:nodeId



## Response
---
  다음 HTTP Status Code 를 사용함





## To-Do
---


1. 1차 구현
    1. Container 조회: 목록 조회, 상세 조회
    2. Container 생성: 생성
    3. Container 시작/종료/삭제: 
2. 2차 구현
    1.