export const API_HOST = "http://mzhn.fun:8080";

export class AuthAPI {
  static getAccessToken(): string | null {
    return localStorage.getItem("accessToken");
  }

  static setAccessToken(token: string): void {
    localStorage.setItem("accessToken", token);
  }

  static getRefreshToken(): string | null {
    return localStorage.getItem("refreshToken");
  }

  static setRefreshToken(token: string): void {
    localStorage.setItem("refreshToken", token);
  }

  static async login(email: string, password: string): Promise<any> {
    const endpoint = `${API_HOST}/api/v1/auth/sign-in`;

    try {
      const data = await this.fetch(endpoint, "POST", { email, password });
      this.setAccessToken(data.data.accessToken);
      this.setRefreshToken(data.data.refreshToken);
      return data;
    } catch (error) {
      console.error("Login failed:", (error as Error).message);
      throw error;
    }
  }

  static async register(
    email: string,
    password: string,
    role: string,
    lastName: string,
    firstName: string,
    middleName: string
  ): Promise<any> {
    const endpoint = `${API_HOST}/api/v1/auth/${role}/sign-up`;
  
    try {
      const data = await this.fetch(endpoint, "POST", {
        email,
        password,
        lastName,
        firstName,
        middleName
      });
      this.setAccessToken(data.data.accessToken);
      this.setRefreshToken(data.data.refreshToken);
      return data;
    } catch (error) {
      console.error("Registration failed:", (error as Error).message);
      throw error;
    }
  }

  static async refreshToken(): Promise<any> {
    const endpoint = `${API_HOST}/api/v1/auth/refresh`;
    const refreshToken = this.getRefreshToken();

    try {
      const data = await this.fetch(endpoint, "POST", { refreshToken });
      this.setAccessToken(data.data.accessToken);
      this.setRefreshToken(data.data.refreshToken);
      return data;
    } catch (error) {
      console.error("Refresh failed:", (error as Error).message);
      throw error;
    }
  }

  static async signOut(): Promise<void> {
    const endpoint = `${API_HOST}/api/v1/auth/sign-out`;
    const accessToken = this.getAccessToken();

    try {
      await this.fetch(endpoint, "POST", null, `Bearer ${accessToken}`);
      console.log("Sign out successful");
    } catch (error) {
      console.error("Sign out failed:", (error as Error).message);
      throw error;
    } finally {
      localStorage.removeItem("accessToken");
      localStorage.removeItem("refreshToken");
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
