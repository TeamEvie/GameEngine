package modules

import dev.kord.core.Kord
import modules.impl.KordPing

/**
 * Holds all modules and is used to instantiate them.
 */
class ModuleManager(
    client: Kord
) {
    val modules = ArrayList<Module>()

    init {
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