import React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { useAuth } from "../contexts/AuthContext";
import WelcomeScreen from "../screens/WelcomeScreen";
import LoginScreen from "../screens/LoginScreen";
import RegisterScreen from "../screens/RegisterScreen";
import HomeScreen from "../screens/HomeScreen";
import DeckListScreen from "../screens/DeckListScreen";
import CreateDeckScreen from "../screens/CreateDeckScreen";
import FlashcardScreen from "../screens/FlashcardScreen";

export type RootStackParamList = {
  Welcome: undefined;
  Login: undefined;
  Register: undefined;
  Home: undefined;
  DeckList: undefined;
  CreateDeck: undefined;
  Flashcards: { deckId: string };
};

const Stack = createNativeStackNavigator<RootStackParamList>();

export default function AppNavigator() {
  const { userId } = useAuth();

  return (
    <Stack.Navigator>
      {userId ? (
        // Protected screens
        <>
          <Stack.Screen
            name="Home"
            component={HomeScreen}
            options={{ title: "FlashGenie" }}
          />
          <Stack.Screen
            name="DeckList"
            component={DeckListScreen}
            options={{ title: "Your Decks" }}
          />
          <Stack.Screen
            name="CreateDeck"
            component={CreateDeckScreen}
            options={{ title: "New Deck" }}
          />
          <Stack.Screen
            name="Flashcards"
            component={FlashcardScreen}
            options={{ title: "Flashcards" }}
          />
        </>
      ) : (
        // Auth flow
        <>
          <Stack.Screen
            name="Welcome"
            component={WelcomeScreen}
            options={{ headerShown: false }}
          />
          <Stack.Screen
            name="Login"
            component={LoginScreen}
            options={{ title: "Log In" }}
          />
          <Stack.Screen
            name="Register"
            component={RegisterScreen}
            options={{ title: "Register" }}
          />
        </>
      )}
    </Stack.Navigator>
  );
}
