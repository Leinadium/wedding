const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

// --- Types ---

export interface Guest {
  name: string;
  phone: string;
}

export interface Product {
  id: string;
  name: string;
  imageUrl: string;
  priceBrl: number;
  purchased: boolean;
}

export interface Payment {
  url: string;
}

export interface ProductListResponse {
  products: Product[];
}

export interface PaymentResponse {
  payment: Payment;
}

export interface ApiError {
  error: string;
}

// --- Helper Handler ---

async function handleResponse<T>(response: Response): Promise<T> {
  if (!response.ok) {
    const errorBody: ApiError = await response
      .json()
      .catch(() => ({ error: "Unknown server error" }));
    throw new Error(
      errorBody.error || `HTTP error! status: ${response.status}`,
    );
  }

  // Handle the 201 "" (empty string) case for the guest POST
  const text = await response.text();
  return text ? JSON.parse(text) : ({} as T);
}

// --- API Functions ---

export const api = {
  async registerGuest(guest: Guest): Promise<void> {
    const res = await fetch(`${API_URL}/v1/guest/`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(guest),
    });
    return handleResponse<void>(res);
  },

  async getProducts(): Promise<ProductListResponse> {
    const res = await fetch(`${API_URL}/v1/product`);
    return handleResponse<ProductListResponse>(res);
  },

  async getPaymentUrl(productId: string): Promise<PaymentResponse> {
    const res = await fetch(`${API_URL}/v1/product/${productId}/payment`);
    return handleResponse<PaymentResponse>(res);
  },
};
