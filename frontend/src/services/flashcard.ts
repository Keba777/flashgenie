import api from "./api";

export interface Flashcard {
    id: string;
    front: string;
    back: string;
    deckId: string;
}

export async function generateFlashcards(deckId: string, content: string): Promise<Flashcard[]> {
    const res = await api.post(`/decks/${deckId}/flashcards`, { content });
    return res.data as Flashcard[];
}

export async function listFlashcards(deckId: string): Promise<Flashcard[]> {
    const res = await api.get(`/decks/${deckId}/flashcards`);
    return res.data as Flashcard[];
}

export async function getFlashcard(id: string): Promise<Flashcard> {
    const res = await api.get(`/flashcards/${id}`);
    return res.data as Flashcard;
}

export async function deleteFlashcard(id: string): Promise<void> {
    await api.delete(`/flashcards/${id}`);
}