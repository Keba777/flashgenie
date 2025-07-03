import api from "./api";

export interface Deck {
    id: string;
    title: string;
    description?: string;
    ownerId: string;
}

export async function createDeck(title: string, description?: string): Promise<Deck> {
    const res = await api.post("/decks", { title, description });
    return res.data as Deck;
}

export async function listDecks(): Promise<Deck[]> {
    const res = await api.get("/decks");
    return res.data as Deck[];
}

export async function updateDeck(id: string, title: string): Promise<void> {
    await api.patch(`/decks/${id}`, { title });
}

export async function deleteDeck(id: string): Promise<void> {
    await api.delete(`/decks/${id}`);
}
