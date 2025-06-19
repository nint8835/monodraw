import ObjC from 'frida-objc-bridge';

if (!ObjC.available) {
    throw new Error('Objective-C runtime is not loaded');
}

function getClassHierarchy(objcObject: ObjC.Object): ObjC.Object[] {
    const hierarchy = [objcObject];
    while (objcObject.$superClass) {
        objcObject = objcObject.$superClass;
        hierarchy.push(objcObject);
    }

    return hierarchy;
}

const classMethods: Record<string, string[]> = {};

for (const [className, objcClass] of Object.entries(ObjC.classes)) {
    if (!className.startsWith('HTModel')) {
        continue;
    }

    const hierarchy = getClassHierarchy(objcClass);

    // Persistable objects inherit from HTModelPersistentObject
    if (hierarchy.some((c) => c.$className === 'HTModelPersistentObject')) {
        classMethods[className] = objcClass.$ownMethods;
        console.log(`Found persistable class: ${objcClass.$className}`);
    }
}

console.log(JSON.stringify(classMethods, null, 2));
