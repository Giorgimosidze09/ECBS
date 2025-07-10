import axios from 'axios';
import { UsersListOutput, CardActivationRequest, CardActivationResponse } from '../types';

//const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080';
const API_URL = 'http://localhost:8080';
//const API_URL = 'http://localhost:5204';

// Define types
export interface CreateUserInput {
    name: string;
    email: string;
    phone: string;
}

export interface AssignCardInput {
    user_id: number;
    card_id: string;
    device_id: number;
    type: string;
}

export interface TopUpInput {
    user_id: number;
    card_id: number;
    balance: number;
    ride_cost:number;
}

export interface ValidateCardInput {
    card_id: number;
}

export interface UsersListInput {
    limit: number;
    offset: number;
}

export interface CardsListInput {
    limit: number;
    offset: number;
}

export interface CardsListOutput {
    id: number;
    user_id: number;
    card_id: string;
    device_id: number;
    total?: number;
}

export interface ChargesListInput {
    limit: number;
    offset: number;
}

export interface BalanceListInput {
    limit: number;
    offset: number;
}


export interface DevicesListInput {
    limit: number;
    offset: number;
}

export interface ChargesListOutput {
    id: number;
    user_id: number;
    amount: number;
    type: string;
    description: string;
    created_at: string;
    total?: number;
}



export interface DevicesListOutput {
    id: number;
    device_id: number;
    location: number;
    active: string;
    description: string;
    InstalledAt: string;
    total?: number;
}


export interface RideCostInput {
ride_cost: number;
}

export const createUser = async (userData: CreateUserInput) => {
    const response = await axios.post(`${API_URL}/users`, userData);
    return response.data;
};

export const assignCard = async (assignmentData: AssignCardInput) => {
    const response = await axios.post(`${API_URL}/cards/assign`, assignmentData);
    return response.data;
};

export const topUpBalance = async (topUpData: TopUpInput) => {
    const response = await axios.post(`${API_URL}/balances/topup`, topUpData);
    return response.data;
};

export const validateCard = async (cardData: ValidateCardInput) => {
    const response = await axios.post(`${API_URL}/cards/validate`, cardData);
    return response.data;
};

export const fetchUsersList = async (userList: UsersListInput) => {
    const response = await axios.post(`${API_URL}/users/list`, userList);
    return response.data;
};

export const fetchCardsList = async (input: CardsListInput) => {
    const response = await axios.post(`${API_URL}/cards/list`, input);
    return response.data;
};


export const fetchChargesList = async (input: ChargesListInput) => {
    const response = await axios.post(`${API_URL}/charges/list`, input);
    return response.data;
};


export const rideCost = async (rideCost: RideCostInput) => {
    const response = await axios.post(`${API_URL}/balances/ride-cost`, rideCost);
    return response.data;
};


export const fetchBalanceList = async (input: BalanceListInput) => {
    const response = await axios.post(`${API_URL}/balances/list`, input);
    return response.data;
};

export const fetchDevicesList = async (input: { limit: number; offset: number }) => {
  const response = await axios.post(`${API_URL}/devices/list`, input);
  return response.data;
};

export const activateCard = async (activationData: CardActivationRequest) => {
  const response = await axios.post(`${API_URL}/cards/activate`, activationData);
  return response.data as CardActivationResponse;
};

export const fetchCustomerSumBalance = async (deviceId: string): Promise<{ total_balance: number }> => {
  const response = await axios.post(`${API_URL}/customer/sum-balance`, { device_id: deviceId });
  return response.data;
};

// --- USERS ---
export const updateUser = async (id: number, userData: any) => {
  const response = await axios.put(`${API_URL}/users/${id}`, userData);
  return response.data;
};

export const deleteUser = async (id: number) => {
  const response = await axios.delete(`${API_URL}/users/${id}`);
  return response.data;
};

export const getUserById = async (id: number) => {
  const response = await axios.get(`${API_URL}/users/${id}`);
  return response.data;
};

// --- CARDS ---
export const updateCard = async (id: number, cardData: any) => {
  const response = await axios.put(`${API_URL}/cards/${id}`, cardData);
  return response.data;
};

export const deleteCard = async (id: number) => {
  const response = await axios.delete(`${API_URL}/cards/${id}`);
  return response.data;
};

export const getCardById = async (id: number) => {
  const response = await axios.get(`${API_URL}/cards/${id}`);
  return response.data;
};

// --- DEVICES ---
export const getDeviceById = async (id: number) => {
  const response = await axios.get(`${API_URL}/devices/${id}`);
  return response.data;
};

export const updateDevice = async (id: number, deviceData: any) => {
  const response = await axios.put(`${API_URL}/devices/${id}`, deviceData);
  return response.data;
};

export const deleteDevice = async (id: number) => {
  const response = await axios.delete(`${API_URL}/devices/${id}`);
  return response.data;
};

export const createDevice = async (deviceData: { device_id: string; location: string }) => {
  const response = await axios.post(`${API_URL}/devices`, deviceData);
  return response.data;
};

export interface LoginRequest {
  username: string;
  password: string;
}
export interface LoginResponse {
  token: string;
  role: string;
}
export interface RegisterRequest {
  username: string;
  password: string;
  role: string;
}
export interface RegisterResponse {
  id: number;
  username: string;
  role: string;
}

export const login = async (data: LoginRequest): Promise<LoginResponse> => {
  const response = await axios.post(`${API_URL}/auth/login`, data);
  return response.data;
};

export const register = async (data: RegisterRequest): Promise<RegisterResponse> => {
  const response = await axios.post(`${API_URL}/auth/register`, data);
  return response.data;
};

axios.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) config.headers['Authorization'] = `Bearer ${token}`;
  return config;
});
