# Figma MCP Server

Figma REST API를 래핑하여 AI 에이전트(LLM)가 사용할 수 있도록 Model Context Protocol (MCP) 서버 형태로 제공하는 Go 기반 프로젝트입니다.

이 프로젝트는 Figma의 거의 모든 API 엔드포인트를 지원하며, 특히 **편집 기능과 사용 편의성을 고도화한 유틸리티 툴**들을 포함하고 있습니다.

## 주요 기능

### 1. 코어 API 래핑
Figma의 공식 REST API를 충실히 래핑하여 제공합니다.
- **Files & Nodes**: 파일 구조 조회, 특정 노드 검색, 이미지/에셋 추출
- **Comments**: 코멘트 조회, 작성, 삭제
- **Projects & Teams**: 팀 프로젝트, 파일, 컴포넌트, 스타일 목록 조회
- **Variables (Design Tokens)**: 로컬/퍼블리시된 변수 조회 및 생성 (Enterprise 플랜 필요)
- **Webhooks**: 웹훅 생성, 조회, 삭제
- **Dev Resources**: 개발 리소스 링크 관리

### 2. 고도화된 유틸리티 툴 (편집/사용성 중심)
단순 API 호출을 넘어, 자주 쓰이는 패턴을 조합한 강력한 유틸리티를 제공합니다.

- `figma_search_nodes`: 파일 내에서 이름으로 특정 노드를 쉽게 검색
- `figma_get_nodes_by_type`: 특정 타입(TEXT, FRAME, COMPONENT 등)의 노드만 일괄 추출
- `figma_export_frames`: 파일 내의 모든 프레임과 컴포넌트를 한 번에 이미지로 추출
- `figma_export_node_as_image`: 특정 노드를 이미지로 렌더링하여 URL 반환
- `figma_upload_image_from_url`: **외부 이미지 URL을 받아 Figma 파일에 직접 업로드** — 이미지를 다운로드하여 Figma에 등록하고, 노드의 IMAGE fill에 바로 사용할 수 있는 `image_ref`를 반환
- `figma_upload_multiple_images_from_urls`: 여러 이미지 URL을 한 번에 Figma에 배치 업로드하여 `image_ref` 목록을 반환
- `figma_fetch_image_from_url`: 외부 이미지 URL을 받아 Base64로 변환 (로컬 처리용)
- `figma_search_text`: 파일 내의 모든 텍스트 노드에서 특정 문자열 검색
- `figma_export_design_tokens`: Figma 변수(Variables)를 CSS Custom Properties 또는 JSON 형태로 즉시 추출
- `figma_get_variable_summary`: 복잡한 변수 컬렉션과 모드를 사람이 읽기 쉬운 요약 형태로 제공

## 설치 및 실행

### 요구 사항
- Go 1.25 이상
- Figma Personal Access Token (PAT)

### 빌드
```bash
git clone https://github.com/hawoond/figma-mcp.git
cd figma-mcp
go build -o bin/figma-mcp ./cmd/figma-mcp/
```

### 실행
환경 변수 `FIGMA_ACCESS_TOKEN`에 Figma에서 발급받은 토큰을 설정한 후 실행합니다.

```bash
export FIGMA_ACCESS_TOKEN="your_figma_personal_access_token"
./bin/figma-mcp
```

## MCP 클라이언트 설정 예시 (Claude Desktop)

`claude_desktop_config.json` 파일에 다음과 같이 추가하여 사용할 수 있습니다.

```json
{
  "mcpServers": {
    "figma": {
      "command": "/path/to/figma-mcp/bin/figma-mcp",
      "env": {
        "FIGMA_ACCESS_TOKEN": "your_figma_personal_access_token"
      }
    }
  }
}
```

## 프로젝트 구조

- `cmd/figma-mcp/`: MCP 서버 메인 애플리케이션 및 툴 핸들러
- `pkg/figma/api/`: Figma REST API 엔드포인트별 래퍼
- `pkg/figma/client/`: HTTP 클라이언트 코어 (인증, 에러 처리)
- `pkg/figma/types/`: Figma API 응답/요청 JSON 구조체 정의
- `pkg/figma/util/`: 노드 탐색, 이미지 처리, 변수 추출 등 고도화된 유틸리티 함수

## 라이선스

MIT License
