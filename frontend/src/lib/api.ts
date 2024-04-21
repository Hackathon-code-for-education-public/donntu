import { AuthAPI } from "./auth";
import { API_HOST } from "./config";

export class API {
  static async request(
    url: string,
    method: "GET" | "POST",
    body?: any
  ): Promise<any> {
    try {
      const response = await this.makeRequest(url, method, body);
      return response;
    } catch (error) {
      if (error instanceof Error && error.message.includes("401")) {
        // If a 401 error occurred, attempt to refresh the token
        const refreshSuccess = await AuthAPI.refreshToken();
        if (refreshSuccess) {
          // If the refresh was successful, retry the original request
          return this.makeRequest(url, method, body);
        } else {
          throw new Error("Refresh token failed, please login again.");
        }
      } else {
        // Re-throw the error if it's not a 401 or after a failed refresh
        throw error;
      }
    }
  }

  // Helper function to perform fetch operations
  private static async makeRequest(
    url: string,
    method: "GET" | "POST",
    body?: any
  ): Promise<any> {
    const accessToken = AuthAPI.getAccessToken();
    const headers: HeadersInit = { "Content-Type": "application/json" };
    if (accessToken) {
      headers["Authorization"] = `Bearer ${accessToken}`;
    }

    const response = await fetch(url, {
      method: method,
      headers: headers,
      body: JSON.stringify(body),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return response.json();
  }
  // Method for uploading files with form-data
  static async uploadPanorama(
    firstLocation: File,
    secondLocation: File,
    universityId: string,
    name: string,
    address: string,
    type: string
  ): Promise<any> {
    const url = `${API_HOST}/api/v1/panoramas`;
    const formData = new FormData();
    formData.append("firstLocation", firstLocation);
    formData.append("secondLocation", secondLocation);
    formData.append("universityId", universityId);
    formData.append("name", name);
    formData.append("address", address);
    formData.append("type", type);

    try {
      // Make the request with the form data
      return await this.requestWithFormData(url, "POST", formData);
    } catch (error) {
      console.error("Upload failed:", error);
      throw error;
    }
  }

  static async createReview(
    universityId: string,
    sentiment: string,
    text: string
  ): Promise<any> {
    const url = `${API_HOST}/api/v1/reviews`;
    const body = {
      universityId,
      sentiment,
      text,
    };

    try {
      // Make the request with the JSON body
      return await this.request(url, "POST", body);
    } catch (error) {
      console.error("Create review failed:", error);
      throw error;
    }
  }

  // Helper method to make a request with FormData
  private static async requestWithFormData(
    url: string,
    method: "POST",
    formData: FormData,
    retry: boolean = true
  ): Promise<any> {
    const accessToken = AuthAPI.getAccessToken();
    const headers: HeadersInit = {};
    if (accessToken) {
      headers["Authorization"] = `Bearer ${accessToken}`;
    }

    const response = await fetch(url, {
      method: method,
      headers: headers,
      body: formData,
    });

    if (response.status === 401 && retry) {
      // Attempt to refresh the token only once
      const refreshSuccess = await AuthAPI.refreshToken();
      if (refreshSuccess) {
        return this.requestWithFormData(url, method, formData, false); // Retry the request without the retry flag
      } else {
        throw new Error("Unable to refresh token");
      }
    }

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return response.json();
  }
}
