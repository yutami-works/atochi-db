// frontend/app/page.tsx

type History = {
  id: number;
  name: string;
  floor_info: string | null;
  note: string | null;
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

export default async function Home() {
  const res = await fetch('http://localhost:8080/api/locations', { cache: 'no-store' });
  const locations: Location[] = await res.json();

  return (
    <main className="min-h-screen p-8 bg-zinc-50 dark:bg-black font-sans">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-3xl font-bold mb-8 text-black dark:text-white">跡地DB</h1>

        {locations.map((location) => (
          <div key={location.id} className="bg-white dark:bg-zinc-900 border border-zinc-200 dark:border-zinc-800 rounded-xl p-6 mb-6 shadow-sm">
            <h2 className="text-2xl font-semibold mb-2 text-black dark:text-white">{location.name}</h2>
            <p className="text-zinc-500 dark:text-zinc-400 mb-6">住所: {location.address}</p>

            <h3 className="text-lg font-medium mb-3 text-black dark:text-zinc-200">履歴:</h3>
            <ul className="space-y-3">
              {location.histories.map((hist) => (
                <li key={hist.id} className="flex flex-col sm:flex-row sm:items-start border-b border-zinc-100 dark:border-zinc-800 pb-2 last:border-0">
                  <span className="font-bold text-zinc-700 dark:text-zinc-300 min-w-[5rem] mb-1 sm:mb-0">
                    {hist.floor_info && hist.floor_info !== '-' ? `[${hist.floor_info}]` : ''}
                  </span>
                  <span className="text-black dark:text-white">
                    {hist.name}
                    {hist.note && (
                      <span className="ml-2 text-sm text-zinc-500 dark:text-zinc-400">({hist.note})</span>
                    )}
                  </span>
                </li>
              ))}
            </ul>
          </div>
        ))}
      </div>
    </main>
  );
}