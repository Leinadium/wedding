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

export interface Attendee {
  id: string;
  name: string;
  isChild: boolean;
  confirmed: boolean | null;
}

export interface InviteResponse {
  id: string;
  phone: string;
  note: string;
  attendees: Attendee[];
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
  async getProducts(): Promise<ProductListResponse> {
    const res = await fetch(`${API_URL}/v1/product`);
    return handleResponse<ProductListResponse>(res);
  },

  async getPaymentUrl(productId: string): Promise<PaymentResponse> {
    const res = await fetch(`${API_URL}/v1/product/${productId}/payment`);
    return handleResponse<PaymentResponse>(res);
  },

  async getInvite(inviteCode: string): Promise<InviteResponse> {
    const res = await fetch(`${API_URL}/v1/invite/${inviteCode}`);
    return handleResponse<InviteResponse>(res);
  },

  async saveInviteNote(code: string, note: string): Promise<void> {
    const res = await fetch(`${API_URL}/v1/invite/${code}/note`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ note: note }),
    });
    return handleResponse<void>(res);
  },

  async saveAttendee(attendee: Attendee): Promise<void> {
    const res = await fetch(`${API_URL}/v1/attendee/${attendee.id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        isChild: attendee.isChild,
        confirmed: attendee.confirmed,
      }),
    });
    return handleResponse<void>(res);
  },
};
