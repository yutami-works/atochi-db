-- 1. locations テーブル（場所のガワと代表情報）
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    created_by VARCHAR(50), -- 初期は管理者名を直書きで想定
    updated_by VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 2. histories テーブル（テナントの変遷）
CREATE TABLE histories (
    id SERIAL PRIMARY KEY,
    location_id INTEGER NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    floor_info VARCHAR(100), -- "1F", "A区画" など。基本はNULL
    note TEXT,
    start_date DATE, -- 開始年月日（不明ならNULL）
    end_date DATE,   -- 終了年月日（不明ならNULL）
    display_order INTEGER NOT NULL DEFAULT 0, -- 数字が大きいほど新しい想定
    image_url TEXT,
    sv_url TEXT,     -- ストリートビューリンク
    evidence_url TEXT, -- 証拠となるブログや記事のリンク
    created_by VARCHAR(50),
    updated_by VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 3. comments テーブル（ユーザーの交流・タレコミ）
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    location_id INTEGER NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    user_id VARCHAR(50),      -- 将来のユーザー管理用（今はNULL許容）
    guest_name VARCHAR(100),  -- 未ログインユーザーの任意入力名
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- おまけ：Next.jsの画面確認用に「千葉鑑定団 八千代店」の初期データを投入
INSERT INTO locations (name, address, latitude, longitude, created_by, updated_by)
VALUES ('千葉鑑定団 八千代店', '千葉県八千代市勝田台南１丁目１８−１', 35.7118, 140.1255, 'admin', 'admin');

-- ※ location_id が 1 になる前提で履歴データを投入（display_orderが大きいほど新しい）
INSERT INTO histories (location_id, name, floor_info, note, display_order, created_by, updated_by)
VALUES
(1, 'ミドリ電化 八千代店', '1-2F', 'エディオンへ移管', 10, 'admin', 'admin'),
(1, 'aprecio 八千代店', '3F', 'ミドリ電化と同居', 10, 'admin', 'admin'),
(1, 'エディオン 八千代店', '1-2F', 'ミドリ電化から店名変更', 20, 'admin', 'admin'),
(1, '千葉鑑定団 八千代店', '-', '現在の店舗', 30, 'admin', 'admin');

-- 初期コメントのテストデータ
INSERT INTO comments (location_id, guest_name, content)
VALUES
(1, '地元民', 'ミドリ電化のオープン時は結構賑わってましたね。'),
(1, '村神', 'ネカフェよく泊まってたけど無くなったんスね');
