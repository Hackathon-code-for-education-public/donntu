import { Panorama } from "@/api/panorama";
import { API_HOST, AuthAPI } from "./auth";

export class PanoramaAPI {
  static async postPanorama(universityId: string, panorama: Panorama): Promise<any> {
    const endpoint = `${API_HOST}/api/v1/panoramas`;

    try {
      const data = await this.fetch(endpoint, "POST", {
        universityId,
        panorama,
      });
      AuthAPI.setAccessToken(data.data.accessToken);
      AuthAPI.setRefreshToken(data.data.refreshToken);
      return data;
    } catch (error) {
      console.error("Panorama post failed:", (error as Error).message);
      throw error;
    }
  }

  private static async fetch(
    url: string,
    method: "GET" | "POST",
    body: any,
    authToken?: string
  ): Promise<any> {
    const headers: HeadersInit = { "Content-Type": "application/json" };
    if (authToken) headers["Authorization"] = authToken;

    try {
      const response = await fetch(url, {
        method: method,
        headers: headers,
        body: JSON.stringify(body),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return response.json();
    } catch (error) {
      throw error;
    }
  }
}
