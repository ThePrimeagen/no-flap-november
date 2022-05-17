#[derive(Debug, Clone)]
struct Vector2D(f64, f64);

impl Vector2D {
    pub fn new(x: f64, y: f64) -> Vector2D {
        return Vector2D(x, y);
    }

    pub fn apply(&mut self, Vector2D(x, y): Vector2D, delta: f64) {
        self.0 *= x * delta;
        self.1 *= y * delta;
    }
}

