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
import { register } from "../services/auth";
import axios from "axios";

type Props = NativeStackScreenProps<RootStackParamList, "Register">;

export default function RegisterScreen({ navigation }: Props) {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const onSubmit = async () => {
    try {
      await register(email, password);
      Alert.alert("Success", "Account created");
      navigation.replace("Login");
    } catch (err: any) {
      let message = "Something went wrong";
      if (axios.isAxiosError(err) && err.response?.data) {
        message =
          (err.response.data as any).error ||
          (err.response.data as any).message ||
          message;
      } else if (err instanceof Error) {
        message = err.message;
      }
      Alert.alert("Registration failed", message);
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
        <Text style={styles.buttonText}>Sign Up</Text>
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
    backgroundColor: colors.secondary,
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
