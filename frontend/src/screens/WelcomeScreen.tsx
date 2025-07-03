import React from "react";
import { View, Text, Pressable, StyleSheet } from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { RootStackParamList } from "../navigation/AppNavigator";
import { colors } from "../themes/colors";

type Props = NativeStackScreenProps<RootStackParamList, "Welcome">;

export default function WelcomeScreen({ navigation }: Props) {
  return (
    <View style={styles.container}>
      <Text style={styles.title}>FlashGenie</Text>
      <Text style={styles.subtitle}>AI‑powered flashcard generator</Text>
      <Pressable style={styles.loginButton} onPress={() => navigation.navigate("Login")}>
        <Text style={styles.loginText}>Log In</Text>
      </Pressable>
      <Pressable onPress={() => navigation.navigate("Register")}>
        <Text style={styles.signUpLink}>Sign Up</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex:1,
    justifyContent:"center",
    alignItems:"center",
    backgroundColor: colors.background,
    padding:24,
  },
  title: {
    fontSize:36,
    fontWeight:"bold",
    color: colors.primary,
    marginBottom:8,
  },
  subtitle: {
    fontSize:16,
    color: colors.textSecondary,
    marginBottom:32,
  },
  loginButton: {
    backgroundColor: colors.primary,
    paddingVertical:12,
    paddingHorizontal:32,
    borderRadius:8,
    marginBottom:16,
  },
  loginText: {
    color: colors.surface,
    fontSize:18,
    fontWeight:"600",
  },
  signUpLink: {
    color: colors.secondary,
    fontSize:16,
  },
});
