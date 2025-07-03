import React, { useEffect, useState } from "react";
import {
  View,
  FlatList,
  Text,
  ActivityIndicator,
  StyleSheet,
} from "react-native";
import Flashcard from "../components/Flashcard";
import { listFlashcards } from "../services/flashcard";
import { colors } from "../themes/colors";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { RootStackParamList } from "../navigation/AppNavigator";

type Props = NativeStackScreenProps<RootStackParamList, "Flashcards">;

export default function FlashcardScreen({ route }: Props) {
  const { deckId } = route.params;
  const [flashcards, setFlashcards] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    listFlashcards(deckId)
      .then(setFlashcards)
      .catch(console.error)
      .finally(() => setLoading(false));
  }, [deckId]);

  if (loading) {
    return (
      <View style={styles.centered}>
        <ActivityIndicator size="large" color={colors.primary} />
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
    backgroundColor: colors.background,
    padding: 16,
  },
  emptyText: {
    textAlign: "center",
    color: colors.textSecondary,
    marginTop: 24,
    fontSize: 16,
  },
});
