import React, { useState } from "react";
import {
  View,
  TextInput,
  Pressable,
  Text,
  StyleSheet,
  Alert,
} from "react-native";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { colors } from "../themes/colors";
import { RootStackParamList } from "../navigation/AppNavigator";
import { login } from "../services/auth";
import axios from "axios";

type Props = NativeStackScreenProps<RootStackParamList, "Login">;

export default function LoginScreen({ navigation }: Props) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const onSubmit = async () => {
    try {
      const { userId } = await login(email, password);
      if (userId) {
        navigation.replace("Home");
      } else {
        Alert.alert("Login failed", "Invalid response from server.");
      }
    } catch (err: any) {
      let message = "Check your credentials or try again.";
      if (axios.isAxiosError(err) && err.response?.data) {
        message =
          (err.response.data as any).error ||
          (err.response.data as any).message ||
          message;
      } else if (err instanceof Error) {
        message = err.message;
      }
      Alert.alert("Login failed", message);
    }
  };

  return (
    <View style={styles.container}>
      <TextInput
        style={styles.input}
        placeholder="Email"
        value={email}
        onChangeText={setEmail}
        keyboardType="email-address"
        autoCapitalize="none"
        placeholderTextColor={colors.textSecondary}
      />
      <TextInput
        style={styles.input}
        placeholder="Password"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
        placeholderTextColor={colors.textSecondary}
      />
      <Pressable style={styles.button} onPress={onSubmit}>
        <Text style={styles.buttonText}>Log In</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
    justifyContent: "center",
    padding: 24,
  },
  input: {
    backgroundColor: colors.surface,
    padding: 12,
    borderRadius: 6,
    marginBottom: 16,
    color: colors.textPrimary,
  },
  button: {
    backgroundColor: colors.primary,
    padding: 12,
    borderRadius: 6,
    alignItems: "center",
  },
  buttonText: {
    color: colors.surface,
    fontSize: 16,
    fontWeight: "600",
  },
});
