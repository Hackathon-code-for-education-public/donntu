import { AuthAPI } from "./auth";
import { API_HOST, API_HOST2 } from "./config";

export const fetcher = async (url: string) => {
  const getHeaders = (token: string | null) => ({
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  });

  try {
    const accessToken = await AuthAPI.getAccessToken();
    let host = API_HOST;
    if (url.startsWith("/api/v1/chats")) {
      host = API_HOST2
    }
    let response = await fetch(host + url , {
      headers: getHeaders(accessToken)
    });

    if (!response.ok) {
      if (response.status === 401) {
        // Refresh the token if there's a 401 error
        const newAccessToken = await AuthAPI.refreshToken();
        response = await fetch(url, {
          headers: getHeaders(newAccessToken)
        });

        if (!response.ok) throw new Error('Failed to fetch data after refreshing the token');
      } else {
        throw new Error(`Request failed with status ${response.status}`);
      }
    }

    return await response.json();
  } catch (error) {    
    throw new Error((error as Error).message);
  }
};

export const fetcherWithData = (url: string) => fetcher(url).then(r => r.data);