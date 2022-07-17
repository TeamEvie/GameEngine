package client

import modules.ModuleManager
import org.javacord.api.DiscordApi
import org.javacord.api.DiscordApiBuilder
import utils.EnvUtils

class GameEngine {
    val client: DiscordApi = DiscordApiBuilder().setToken(EnvUtils.getEnv("DISCORD_TOKEN")).login().get()
    val modules = ModuleManager()

    init {
        modules.init(client)
    }
}