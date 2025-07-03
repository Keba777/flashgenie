import React, { useEffect, useState } from "react";
import {
  View,
  FlatList,
  Text,
  ActivityIndicator,
  StyleSheet,
} from "react-native";
import Flashcard from "../components/Flashcard";
import { listFlashcards } from "../services/api";

export default function FlashcardScreen() {
  const [flashcards, setFlashcards] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  // TODO: Replace 'your-deck-id' with the actual deckId you want to fetch
  useEffect(() => {
    const deckId = "your-deck-id";
    listFlashcards(deckId)
      .then((data) => setFlashcards(data))
      .catch(console.error)
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return (
      <View style={styles.centered}>
        <ActivityIndicator size="large" color="#4F46E5" />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      {flashcards.length === 0 ? (
        <Text style={styles.emptyText}>No flashcards available.</Text>
      ) : (
        <FlatList
          data={flashcards}
          keyExtractor={(item) => item.id}
          renderItem={({ item }) => (
            <Flashcard front={item.front} back={item.back} />
          )}
          contentContainerStyle={{ paddingBottom: 32 }}
        />
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  centered: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
  },
  container: {
    flex: 1,
    backgroundColor: "#F3F4F6", // background
    padding: 16, // p-4
  },
  emptyText: {
    textAlign: "center",
    color: "#6B7280", // text-secondary
    marginTop: 24, // mt-6
    fontSize: 16,
  },
});
