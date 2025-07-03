import React, { useState } from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";

interface FlashcardProps {
  front: string;
  back: string;
}

export default function Flashcard({ front, back }: FlashcardProps) {
  const [flipped, setFlipped] = useState(false);

  return (
    <Pressable style={styles.card} onPress={() => setFlipped(!flipped)}>
      <Text style={styles.textPrimary}>{flipped ? back : front}</Text>
      <Text style={styles.textSecondary}>
        Tap card to {flipped ? "show question" : "show answer"}
      </Text>
    </Pressable>
  );
}

const styles = StyleSheet.create({
  card: {
    backgroundColor: "#FFFFFF", // surface
    borderRadius: 12, // rounded-lg ~12px
    padding: 24, // p-6 (6*4=24)
    marginBottom: 16, // mb-4 (4*4=16)
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 3, // Android shadow
  },
  textPrimary: {
    color: "#111827", // text-primary
    fontSize: 18, // text-lg ~18px
    fontWeight: "500", // font-medium
  },
  textSecondary: {
    color: "#6B7280", // text-secondary
    marginTop: 8, // mt-2 (2*4=8)
    fontSize: 14, // text-sm ~14px
    fontStyle: "italic",
    textAlign: "center",
  },
});
