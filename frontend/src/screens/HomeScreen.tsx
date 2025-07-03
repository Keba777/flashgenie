import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { colors } from "../themes/colors";
import { RootStackParamList } from "../navigation/AppNavigator";

type Props = NativeStackScreenProps<RootStackParamList, "Home">;

export default function HomeScreen({ navigation }: Props) {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>Welcome, Student!</Text>
      <Pressable
        style={[styles.button, { backgroundColor: colors.primary }]}
        onPress={() => navigation.navigate("DeckList")}
      >
        <Text style={styles.btnText}>My Decks</Text>
      </Pressable>
      <Pressable
        style={[styles.button, { backgroundColor: colors.secondary }]}
        onPress={() => navigation.navigate("CreateDeck")}
      >
        <Text style={styles.btnText}>New Deck</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex:1,
    backgroundColor: colors.background,
    justifyContent:"center",
    alignItems:"center",
    padding:24,
  },
  title: {
    fontSize:28,
    color: colors.textPrimary,
    marginBottom:24,
  },
  button: {
    paddingVertical:12,
    paddingHorizontal:32,
    borderRadius:8,
    marginVertical:8,
  },
  btnText: {
    color: colors.surface,
    fontSize:16,
    fontWeight:"600",
  },
});
