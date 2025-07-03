import React, { useState } from "react";
import {
  View,
  TextInput,
  Pressable,
  Text,
  StyleSheet,
  Alert,
} from "react-native";
import { createDeck } from "../services/deck";
import { colors } from "../themes/colors";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { RootStackParamList } from "../navigation/AppNavigator";

type Props = NativeStackScreenProps<RootStackParamList, "CreateDeck">;

export default function CreateDeckScreen({ navigation }: Props) {
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");

  const onSubmit = async () => {
    try {
      await createDeck(title, desc);
      navigation.goBack();
    } catch {
      Alert.alert("Failed to create deck");
    }
  };

  return (
    <View style={styles.container}>
      <TextInput
        style={styles.input}
        placeholder="Deck Title"
        value={title}
        onChangeText={setTitle}
        placeholderTextColor={colors.textSecondary}
      />
      <TextInput
        style={[styles.input, { height: 100 }]}
        placeholder="Description"
        value={desc}
        onChangeText={setDesc}
        multiline
        placeholderTextColor={colors.textSecondary}
      />
      <Pressable style={styles.button} onPress={onSubmit}>
        <Text style={styles.btnText}>Create</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.background,
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
  btnText: {
    color: colors.surface,
    fontSize: 16,
    fontWeight: "600",
  },
});
