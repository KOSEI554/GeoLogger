import { useEffect, useRef, useState } from "react";
import L from "leaflet";
import "leaflet/dist/leaflet.css";
import "./App.css";

interface Location {
  lat: number;
  lng: number;
}

function App() {
  const mapRef = useRef<HTMLDivElement>(null);
  const [map, setMap] = useState<L.Map | null>(null);
  const [positions, setPositions] = useState<Location[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const markersRef = useRef<L.Marker[]>([]);

  const sendLocation = async (lat: number, lng: number) => {
    setIsLoading(true);
    try {
      await fetch("http://localhost:8080/api/locations", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ lat, lng }),
      });

      // 位置情報を送信した後で再取得する
      await fetchLocations();
    } catch (error) {
      console.error("位置情報の送信に失敗しました:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const getCurrentLocation = () => {
    setIsLoading(true);
    // ブラウザの位置情報APIを使用して現在地を取得
    navigator.geolocation.getCurrentPosition(
      (pos) => {
        const { latitude, longitude } = pos.coords;
        sendLocation(latitude, longitude);
      },
      (error) => {
        console.error("位置情報の取得に失敗しました:", error);
        setIsLoading(false);
      }
    );
  };

  const fetchLocations = async () => {
    if (!map) return;

    try {
      const res = await fetch("http://localhost:8080/api/locations");
      const data: Location[] = await res.json();

      // 既存のマーカーをクリア
      markersRef.current.forEach((marker) => map.removeLayer(marker));
      markersRef.current = [];

      // 新しいマーカーを追加
      const newMarkers = data.map(({ lat, lng }) => {
        const marker = L.marker([lat, lng]).addTo(map);
        return marker;
      });

      markersRef.current = newMarkers;
      setPositions(data);

      // マーカーがある場合は適切な表示範囲に調整
      if (data.length > 0) {
        const bounds = L.latLngBounds(data.map((loc) => [loc.lat, loc.lng]));
        map.fitBounds(bounds, { padding: [50, 50] });
      }
    } catch (error) {
      console.error("位置情報の取得に失敗しました:", error);
    }
  };

  useEffect(() => {
    if (!mapRef.current) return;

    // 地図の初期化
    const mapInstance = L.map(mapRef.current).setView([35.6895, 139.6917], 13);
    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution: "&copy; OpenStreetMap contributors",
    }).addTo(mapInstance);

    setMap(mapInstance);

    // クリーンアップ関数
    return () => {
      mapInstance.remove();
    };
  }, []);

  useEffect(() => {
    if (map) {
      fetchLocations();
    }
  }, [map]);

  return (
    <div className="app-container">
      <h1>位置情報ロガー</h1>
      <div className="controls">
        <button onClick={getCurrentLocation} disabled={isLoading}>
          {isLoading ? "処理中..." : "📍 現在地を送信"}
        </button>
      </div>
      <div
        ref={mapRef}
        className="map-container"
        style={{ height: "700px", marginTop: "20px", borderRadius: "8px" }}
      ></div>
      <div className="info">
        <p>記録された位置: {positions.length}箇所</p>
      </div>
    </div>
  );
}

export default App;
