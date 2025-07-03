import React, { useEffect, useState } from "react";
import {
  View,
  Text,
  FlatList,
  Pressable,
  StyleSheet,
  ActivityIndicator,
} from "react-native";
import { listDecks } from "../services/deck";
import { colors } from "../themes/colors";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { RootStackParamList } from "../navigation/AppNavigator";

type Props = NativeStackScreenProps<RootStackParamList, "DeckList">;

type Deck = {
  id: string;
  title: string;
  description?: string;
};

export default function DeckListScreen({ navigation }: Props) {
  const [decks, setDecks] = useState<Deck[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchDecks = async () => {
      try {
        const data = await listDecks();
        console.log("📦 Decks from API:", data);

        // ✅ map backend keys to match frontend expectations
        const mappedDecks: Deck[] = data.map((item: any) => ({
          id: item.ID,
          title: item.Title,
          description: item.Description,
        }));

        setDecks(mappedDecks);
      } catch (error) {
        console.error("❌ Failed to fetch decks:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchDecks();
  }, []);

  if (loading) {
    return (
      <View style={styles.centered}>
        <ActivityIndicator size="large" color={colors.primary} />
      </View>
    );
  }

  return (
    <View style={styles.container}>
      <FlatList
        data={decks}
        keyExtractor={(item, index) =>
          item?.id ? String(item.id) : `deck-${index}`
        }
        renderItem={({ item }) => (
          <Pressable
            style={styles.deck}
            onPress={() =>
              navigation.navigate("Flashcards", { deckId: item.id })
            }
          >
            <Text style={styles.deckTitle}>{item.title}</Text>
            <Text style={styles.deckDesc}>
              {item.description || "No description"}
            </Text>
          </Pressable>
        )}
        contentContainerStyle={{ paddingBottom: 16 }}
      />
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
  deck: {
    backgroundColor: colors.surface,
    padding: 16,
    borderRadius: 8,
    marginBottom: 12,
    elevation: 2,
  },
  deckTitle: {
    fontSize: 18,
    color: colors.textPrimary,
    fontWeight: "500",
  },
  deckDesc: {
    fontSize: 14,
    color: colors.textSecondary,
    marginTop: 4,
  },
});
