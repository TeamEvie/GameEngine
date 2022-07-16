import dev.kord.core.Kord
import dev.kord.gateway.Intent
import dev.kord.gateway.PrivilegedIntent
import modules.ModuleManager

suspend fun main() {
    val kord = Kord(System.getenv("DISCORD_TOKEN"))

    ModuleManager(kord)

    kord.login {
        @OptIn(PrivilegedIntent::class)
        intents += Intent.MessageContent
    }
}