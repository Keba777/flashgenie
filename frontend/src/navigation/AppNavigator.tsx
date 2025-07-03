import React from "react";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import HomeScreen from "../screens/HomeScreen";
import FlashcardScreen from "../screens/FlashcardScreen";

const Stack = createNativeStackNavigator();

export default function AppNavigator() {
  return (
    <NavigationContainer>
      <Stack.Navigator
        initialRouteName="Home"
        screenOptions={{ headerStyle: { backgroundColor: "#4F46E5" }, headerTintColor: "#fff" }}
      >
        <Stack.Screen name="Home" component={HomeScreen} options={{ title: "FlashGenie" }} />
        <Stack.Screen name="Flashcards" component={FlashcardScreen} options={{ title: "Your Flashcards" }} />
      </Stack.Navigator>
    </NavigationContainer>
  );
}
