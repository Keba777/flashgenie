import axios from "axios";

// Replace with the IP you found, and your backend port
const DEVICE_IP = "192.168.100.89";
const API_BASE_URL = `http://${DEVICE_IP}:8080/api`;

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 5000,
  // headers: { Authorization: `Bearer ${token}` },
});

// Auth
export async function register(email: string, password: string) {
  const res = await api.post("/auth/register", { email, password });
  return res.data;
}

export async function login(email: string, password: string) {
  const res = await api.post("/auth/login", { email, password });
  return res.data.token; // assuming { token: "…" }
}

// Decks
export async function createDeck(title: string, description?: string) {
  const res = await api.post("/decks", { title, description });
  return res.data; // newly created deck object
}

export async function listDecks() {
  const res = await api.get("/decks");
  return res.data; // array of decks
}

export async function updateDeck(id: string, title: string) {
  await api.patch(`/decks/${id}`, { title });
}

export async function deleteDeck(id: string) {
  await api.delete(`/decks/${id}`);
}

// Flashcards
export async function generateFlashcards(deckId: string, content: string) {
  const res = await api.post(`/decks/${deckId}/flashcards`, { content });
  return res.data; // array of generated flashcards
}

export async function listFlashcards(deckId: string) {
  const res = await api.get(`/decks/${deckId}/flashcards`);
  return res.data; // array of flashcards
}

export async function getFlashcard(id: string) {
  const res = await api.get(`/flashcards/${id}`);
  return res.data; // single flashcard
}

export async function deleteFlashcard(id: string) {
  await api.delete(`/flashcards/${id}`);
}

// Test OpenAI
export async function testOpenAIKey() {
  const res = await api.get("/test-openai");
  return res.data; // { message, available_models_count }
}
