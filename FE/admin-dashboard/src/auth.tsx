import React, { createContext, useContext, useState } from 'react';
import { jwtDecode } from 'jwt-decode';

interface AuthContextType {
  token: string | null;
  role: string | null;
  deviceId: string | null;
  login: (token: string, role: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [token, setToken] = useState<string | null>(() => localStorage.getItem('token'));
  const [role, setRole] = useState<string | null>(() => localStorage.getItem('role'));
  const [deviceId, setDeviceId] = useState<string | null>(() => localStorage.getItem('device_id'));

  const login = (token: string, role: string) => {
    setToken(token);
    setRole(role);
    localStorage.setItem('token', token);
    localStorage.setItem('role', role);
    try {
      const decoded: any = jwtDecode(token);
      const deviceId = decoded.device_id || null;
      setDeviceId(deviceId);
      if (deviceId) localStorage.setItem('device_id', deviceId);
      else localStorage.removeItem('device_id');
    } catch {
      setDeviceId(null);
      localStorage.removeItem('device_id');
    }
  };

  const logout = () => {
    setToken(null);
    setRole(null);
    setDeviceId(null);
    localStorage.removeItem('token');
    localStorage.removeItem('role');
    localStorage.removeItem('device_id');
  };

  return (
    <AuthContext.Provider value={{ token, role, deviceId, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error('useAuth must be used within AuthProvider');
  return ctx;
}; 