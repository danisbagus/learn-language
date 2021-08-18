<?php

class Person {
    private $name;
    private $age;

    function __construct($name, $age) {
        $this->name = $name;
        $this->age = $age;
    }

    public function setName ($name) {
        $this->name = $name;
    }

    public function getName () {
        echo "Get Name: " . $this->name . "\n";
    }
}

    $Person1 = new Person("Hera", 20);
    $Person1->getName();
    $Person1->setName("Hero");
    $Person1->getName();

?>