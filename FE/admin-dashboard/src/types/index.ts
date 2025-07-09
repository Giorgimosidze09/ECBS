export interface User {
    id: string;
    name: string;
    email: string;
    // Add other user fields as necessary
}

export interface Card {
    id: number;
    user_id: number;
    card_id: string;
    device_id: number;
    type?: string;
    active?: boolean;
    assigned_at?: string;
    // Add other card fields as necessary
}

export interface Balance {
    userId: string;
    amount: number;
}

export interface ValidateCardResponse {
    isValid: boolean;
    message: string;
}

export interface TopUpBalanceRequest {
    userId: string;
    amount: number;
}

export interface AssignCardInput {
    user_id: number;
    card_id: string;
    device_id: number;
    type: string; // 'balance' or 'activation'
}

export interface AssignCardRequest {
    userId: string;
    cardNumber: string;
    type: string; // 'balance' or 'activation'
}

export interface CreateUserRequest {
    name: string;
    email: string;
    // Add other fields as necessary
}

export interface UsersListOutput {
    id: number;
    name: string;
    email: string;
    phone: string;
    card_count: number;
    total_balance: number;
    total?: number;
}

export interface CardActivationRequest {
    cardId: number;
    activationStart: string; // ISO date string
    activationEnd: string;   // ISO date string
}

export interface CardActivationResponse {
    id: number;
    cardId: number;
    activationStart: string;
    activationEnd: string;
}

export interface Device {
    id: number;
    device_id: string;
    location: string;
    installed_at: string;
    active: boolean;
}