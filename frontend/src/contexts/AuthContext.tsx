import React, { createContext, useContext, useState, useEffect, ReactNode } from "react";
import * as authService from "../services/auth";

interface AuthContextType {
  userId: string | null;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [userId, setUserId] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      const id = await authService.getCurrentUserId();
      setUserId(id);
    })();
  }, []);

  const handleLogin = async (email: string, password: string) => {
    const { userId: id } = await authService.login(email, password);
    setUserId(id);
  };

  const handleRegister = async (email: string, password: string) => {
    await authService.register(email, password);
    const { userId: id } = await authService.login(email, password);
    setUserId(id);
  };

  const handleLogout = async () => {
    await authService.logout();
    setUserId(null);
  };

  return (
    <AuthContext.Provider value={{ userId, login: handleLogin, register: handleRegister, logout: handleLogout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within AuthProvider");
  return context;
}