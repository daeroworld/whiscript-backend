# whiscript-backend

## 🧑‍💻: Intro

Go 기반 백엔드로, 음성-텍스트 변환, 자막 편집, 미디어 처리를 위한 CQRS/이벤트 기반 마이크로서비스를 제공합니다.
이 백엔드는 모듈식으로 설계되었으며, 각 핵심 서비스(text, voice, shared)는 독립적인 개발과 버전 관리를 위해 별도의 Git 서브모듈로 구성되어 있습니다.

### ❓ Problem

영상 콘텐츠 제작자들은 대용량 미디어 파일을 처리하고, 자막을 생성하며, 실시간 편집을 지원하는 제작자 친화적인 편집기가 필요합니다.

### ❗ Idea

CQRS, 이벤트 기반 아키텍처, 마이크로서비스를 활용하여 텍스트 및 음성 처리를 위한 백엔드를 설계합니다.
각 서비스를 독립적으로 유지보수하고 업그레이드할 수 있도록 서브모듈을 사용합니다.


### 💯 Solution

- **CQRS/이벤트 기반**: 명령/쿼리 분리 및 이벤트 버스로 확장성 확보
- **청크 업로드**: 파일을 분할하여 효율적으로 대용량 파일 처리
- **Whisper 통합**: Whisper CLI/API를 호출하여 높은 정확도의 음성-텍스트 변환
- **gRPC 게이트웨이**: 자막 편집 및 처리에 대한 실시간 업데이트
- **서브모듈 아키텍처**: 각 서비스(text, voice, shared)는 독립적인 Git 서브모듈

## 🧱: Structure

```
backend/
├── monolithic/      # 올인원 Go 서비스 (레거시 또는 참고용)
├── shared/          # 공유 Go 코드 및 인터페이스 (submodule)
├── text/
│   ├── reader/      # 자막 쿼리 마이크로서비스 (submodule)
│   ├── writer/      # 자막 커맨드 마이크로서비스 (submodule)
│   └── bus/         # 텍스트 서비스용 이벤트 버스 (submodule)
├── voice/
│   ├── reader/      # 음성 쿼리 마이크로서비스 (submodule)
│   ├── writer/      # 음성 커맨드 마이크로서비스 (submodule)
│   └── bus/         # 음성 서비스용 이벤트 버스 (submodule)
├── postgresql/      # DB 설정 및 Dockerfile
├── .gitmodules      # 모든 서브모듈과 URL 선언
└── ...
```


## 🚀: Build & Run
```bash
git clone --recursive 
cd backend

# DB 시작 (docker-compose 사용 시)
docker-compose up -d postgres

# 올인원 서비스 (개발/테스트용)
cd monolithic
go mod download
go run ./cmd/main.go

# 또는 개별 마이크로서비스 실행 (text/reader, text/writer 등)
cd ../text/reader
go mod download
go run .
```

## 🔥: Accomplishments

- Go 기반의 모듈식 CQRS/이벤트 기반 백엔드
- 실시간 자막 및 음성 처리
- 원활한 청크 업로드 및 재조립
- Whisper CLI/API 통합
- 확장 가능하고 독립적인 서비스 개발을 위한 서브모듈 아키텍처

## ✅: Implementation

- [x] CQRS/이벤트 기반 구조
- [x] 청크 파일 업로드 및 재조립
- [x] Whisper 통합
- [x] 실시간 업데이트를 위한 WebSocket 게이트웨이
- [x] 모든 핵심 서비스의 서브모듈 관리

## 🛠️: Technologies Used

- Go (Golang)
- PostgreSQL
- FFmpeg
- Whisper (OpenAI)
- Docker

## 📚: Libraries Used

- gin
- gorm
- go-uuid
- gRPC, protobuf
