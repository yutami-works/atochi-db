// frontend/app/page.tsx

type History = {
  id: number;
  name: string;
  floor_info: string | null;
  note: string | null;
  start_date: string | null;
  end_date: string | null;
  image_url: string | null;
  display_order: number;
};

type Location = {
  id: number;
  name: string;
  address: string;
  latitude: number;
  longitude: number;
  histories: History[];
};

// 日付のフォーマット処理（無い場合は「不明」を返す）
const formatDate = (dateString: string | null) => {
  if (!dateString) return '不明';
  // DBのタイムスタンプ文字列（2005-01-01T...）から日付部分だけ切り取る
  return dateString.split('T')[0];
};

export default async function Home() {
  const res = await fetch('http://localhost:8080/api/locations', { cache: 'no-store' });
  const locations: Location[] = await res.json();

  return (
    <main className="min-h-screen p-8 bg-zinc-50 dark:bg-black font-sans">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-3xl font-bold mb-12 text-black dark:text-white">跡地DB</h1>

        {locations.map((location) => (
          <div key={location.id} className="mb-16">

            {/* 場所のヘッダー */}
            <div className="mb-6 border-b border-zinc-200 dark:border-zinc-800 pb-4">
              <h2 className="text-2xl font-bold text-black dark:text-white flex items-baseline">
                <span className="text-zinc-400 text-lg mr-2 font-mono">#{location.id}</span>
                {location.name}
              </h2>
              <p className="text-zinc-500 dark:text-zinc-400 mt-2">📍 {location.address}</p>
            </div>

            {/* 履歴のリスト（カード型のまとまり） */}
            <div className="space-y-4">
              {location.histories.map((hist) => (
                <div key={hist.id} className="flex flex-col sm:flex-row gap-5 bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl p-4 shadow-sm hover:shadow-md transition-shadow">

                  {/* 画像エリア（左側） */}
                  <div className="w-full sm:w-48 h-32 bg-zinc-100 dark:bg-zinc-800/50 rounded-lg flex-shrink-0 flex items-center justify-center overflow-hidden border border-zinc-100 dark:border-zinc-800">
                    {hist.image_url ? (
                      <img src={hist.image_url} alt={hist.name} className="w-full h-full object-cover" />
                    ) : (
                      <div className="flex flex-col items-center text-zinc-400">
                        <span className="text-2xl mb-1">📷</span>
                        <span className="text-xs font-semibold">No Image</span>
                      </div>
                    )}
                  </div>

                  {/* テキスト情報エリア（右側） */}
                  <div className="flex flex-col justify-center flex-grow py-1">
                    <h3 className="text-xl font-bold text-black dark:text-white flex items-center flex-wrap gap-2 mb-2">
                      {hist.name}
                      {hist.floor_info && hist.floor_info !== '-' && (
                        <span className="text-xs bg-zinc-200 dark:bg-zinc-800 text-zinc-700 dark:text-zinc-300 px-2 py-1 rounded font-medium">
                          {hist.floor_info}
                        </span>
                      )}
                      <span className="text-xs text-zinc-400 font-mono ml-auto">ID:{hist.id}</span>
                    </h3>

                    <p className="text-sm font-mono text-zinc-600 dark:text-zinc-400 mb-3 bg-zinc-50 dark:bg-black inline-block px-2 py-1 rounded border border-zinc-100 dark:border-zinc-800 self-start">
                      🗓️ {formatDate(hist.start_date)} ～ {formatDate(hist.end_date)}
                    </p>

                    {hist.note && (
                      <p className="text-zinc-700 dark:text-zinc-300 text-sm leading-relaxed">
                        {hist.note}
                      </p>
                    )}
                  </div>

                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
    </main>
  );
}