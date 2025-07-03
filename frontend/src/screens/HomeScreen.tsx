import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";

export default function HomeScreen({ navigation }: any) {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Welcome to FlashGenie</Text>
      <Text style={styles.subtitle}>
        AI-powered flashcard generator for smarter learning.
      </Text>

      <Pressable
        onPress={() => navigation.navigate("Flashcards")}
        style={styles.button}
      >
        <Text style={styles.buttonText}>View Flashcards</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#F3F4F6", // background
    justifyContent: "center",
    alignItems: "center",
    paddingHorizontal: 24, // px-6
  },
  title: {
    color: "#111827", // text-primary
    fontSize: 32, // text-4xl
    fontWeight: "bold",
    marginBottom: 24, // mb-6
  },
  subtitle: {
    color: "#6B7280", // text-secondary
    textAlign: "center",
    marginBottom: 48, // mb-12
    fontSize: 16,
  },
  button: {
    backgroundColor: "#4F46E5", // primary
    paddingHorizontal: 24, // px-6
    paddingVertical: 12, // py-3
    borderRadius: 8, // rounded-md
  },
  buttonText: {
    color: "#FFFFFF",
    fontSize: 18,
    fontWeight: "600", // font-semibold
  },
});
