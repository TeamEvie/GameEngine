package modules

import modules.impl.KordPing
import org.javacord.api.DiscordApi

/** Holds all modules and is used to instantiate them. */
class ModuleManager() {
    private val modules = ArrayList<Module>()

    fun init(client: DiscordApi) {
        modules.add(KordPing(client))
    }

    /**
     * Called when the bot is shutting down.
     * @return number of modules that shutdown.
     */
    fun shutdown(): Int {
        for (module in modules) module.shutdown()
        return modules.size
    }
}
