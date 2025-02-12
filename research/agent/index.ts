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

const persistableClasses: ObjC.Object[] = [];

for (const [className, objcClass] of Object.entries(ObjC.classes)) {
    if (!className.startsWith('HTModel')) {
        continue;
    }

    const hierarchy = getClassHierarchy(objcClass);

    // Persistable objects inherit from HTModelPersistentObject
    if (hierarchy.some((c) => c.$className === 'HTModelPersistentObject')) {
        persistableClasses.push(objcClass);
        console.log(`Found persistable class: ${objcClass.$className}`);
    }
}
