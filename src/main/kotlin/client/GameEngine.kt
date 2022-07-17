package client

import modules.ModuleManager
import org.javacord.api.DiscordApi
import org.javacord.api.DiscordApiBuilder
import redis.clients.jedis.JedisPooled
import utils.EnvUtils

class GameEngine {
    val pool = JedisPooled(EnvUtils.getEnv("REDIS_HOST"), Integer.parseInt(EnvUtils.getEnv("REDIS_PORT")))
    val discord: DiscordApi = DiscordApiBuilder().setToken(EnvUtils.getEnv("DISCORD_TOKEN")).login().get()
    val modules = ModuleManager()

    init {
        modules.init(discord)
    }
}