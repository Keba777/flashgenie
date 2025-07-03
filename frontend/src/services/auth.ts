import api from "./api";
import { jwtDecode } from "jwt-decode";
import AsyncStorage from "@react-native-async-storage/async-storage";

interface JwtPayload {
    sub: string;
    exp: number;
    iat: number;
}

export async function register(email: string, password: string) {
    const res = await api.post("/auth/register", { email, password });
    return res.data;
}

export async function login(email: string, password: string) {
    const res = await api.post("/auth/login", { email, password });
    const { token } = res.data;
    await AsyncStorage.setItem("token", token);
    const payload = jwtDecode<JwtPayload>(token);
    return { token, userId: payload.sub };
}

export async function logout() {
    await AsyncStorage.removeItem("token");
}

export async function getToken(): Promise<string | null> {
    return AsyncStorage.getItem("token");
}

export async function getCurrentUserId(): Promise<string | null> {
    const token = await getToken();
    if (!token) return null;
    const payload = jwtDecode<JwtPayload>(token);
    return payload.sub;
}
