import dev.kord.core.Kord
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import modules.ModuleManager
import utils.EnvUtils

suspend fun main() {
    val kord = Kord(EnvUtils.getEnv("DISCORD_TOKEN"))

    ModuleManager(kord)

    kord.login {
        @OptIn(PrivilegedIntent::class) //
        intents += Intent.MessageContent
    }
}
