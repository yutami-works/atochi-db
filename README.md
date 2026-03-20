# atochi-db
仮称：跡地DB

## 開発環境の起動方法

### 1. Docker (PostgreSQL)

プロジェクトルートで実行:

```bash
docker compose up -d
```

→ PostgreSQL が `localhost:5432` で起動

### 2. Go バックエンド

```bash
cd backend
go run main.go
```

### 3. Next.js フロントエンド

```bash
cd frontend
npm run dev
```
→ `http://localhost:3000` でアクセス

> **注意:** バックエンド・フロントエンドは Docker 化されていない。ローカルで直接起動。
